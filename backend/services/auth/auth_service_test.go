package authservice

import (
	"errors"
	"testing"
	"time"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	tokenservice "github.com/Akif-jpg/MyHobieMMORPGGame/services/token"
)

// --- Mock User Repository ---

type mockUserRepo struct {
	users map[string]*models.User // keyed by email
}

func newMockUserRepo(users ...*models.User) *mockUserRepo {
	m := &mockUserRepo{users: make(map[string]*models.User)}
	for _, u := range users {
		m.users[u.Email] = u
	}
	return m
}

func (m *mockUserRepo) GetUserByEmail(email string) (*models.User, error) {
	u, ok := m.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// --- Test Helpers ---

func newTestAuthService(repo UserRepository) *AuthService {
	return NewAuthService(repo, "test-secret", "game-server", time.Minute*5)
}

// --- Tests ---

func TestLogin_Success(t *testing.T) {
	user := &models.User{
		ID:       1,
		Username: "akif",
		Email:    "akif@game.com",
		Role:     "player",
		Password: "correct-password",
	}

	svc := newTestAuthService(newMockUserRepo(user))
	token, err := svc.Login(user.Email, user.Password)
	if err != nil {
		t.Fatalf("expected successful login, got error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	user := &models.User{
		ID:       2,
		Username: "akif",
		Email:    "akif@game.com",
		Role:     "player",
		Password: "correct-password",
	}

	svc := newTestAuthService(newMockUserRepo(user))
	_, err := svc.Login(user.Email, "wrong-password")
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}
	if err != ErrInvalidCredentials {
		t.Errorf("expected ErrInvalidCredentials, got: %v", err)
	}
}

func TestLogin_UserNotFound(t *testing.T) {
	svc := newTestAuthService(newMockUserRepo()) // empty repo
	_, err := svc.Login("nobody@game.com", "any-password")
	if err == nil {
		t.Fatal("expected error for unknown user, got nil")
	}
	if err != ErrUserNotFound {
		t.Errorf("expected ErrUserNotFound, got: %v", err)
	}
}

// TestLogin_TokenContainsRole verifies that the generated token encodes the correct role.
func TestLogin_TokenContainsRole(t *testing.T) {
	user := &models.User{
		ID:       3,
		Username: "admin",
		Email:    "admin@game.com",
		Role:     "admin",
		Password: "adminpass",
	}

	svc := newTestAuthService(newMockUserRepo(user))
	tokenStr, err := svc.Login(user.Email, user.Password)
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	parsed, err := tokenservice.ValidateToken(tokenStr, "test-secret")
	if err != nil {
		t.Fatalf("token validation failed: %v", err)
	}
	if parsed.Body.UserRole != user.Role {
		t.Errorf("expected role %q in token, got %q", user.Role, parsed.Body.UserRole)
	}
}

// TestLogin_TokenContainsUsername verifies that the generated token encodes the correct username.
func TestLogin_TokenContainsUsername(t *testing.T) {
	user := &models.User{
		ID:       4,
		Username: "player1",
		Email:    "player1@game.com",
		Role:     "player",
		Password: "pass",
	}

	svc := newTestAuthService(newMockUserRepo(user))
	tokenStr, err := svc.Login(user.Email, user.Password)
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	parsed, err := tokenservice.ValidateToken(tokenStr, "test-secret")
	if err != nil {
		t.Fatalf("token validation failed: %v", err)
	}
	if parsed.Body.Username != user.Username {
		t.Errorf("expected username %q in token, got %q", user.Username, parsed.Body.Username)
	}
}
