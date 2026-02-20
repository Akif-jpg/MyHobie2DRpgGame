// Package authservice provides login and token generation for game users.
package authservice

import (
	"errors"
	"time"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	tokenservice "github.com/Akif-jpg/MyHobieMMORPGGame/services/token"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)

// UserRepository is the subset of the full UserRepositoryInterface that AuthService needs.
// This allows easy mocking in tests without importing the repositories package.
type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}

// AuthService handles user authentication and issues JWT-style tokens.
type AuthService struct {
	userRepo  UserRepository
	jwtSecret string
	producer  string
	expiry    time.Duration
}

// NewAuthService creates a new AuthService with the given dependencies.
func NewAuthService(userRepo UserRepository, jwtSecret, producer string, expiry time.Duration) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		producer:  producer,
		expiry:    expiry,
	}
}

// Login authenticates a user by email and password.
// On success it returns a signed token string.
// Password comparison is plain-text for now; swap for bcrypt when User.Password stores hashes.
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", ErrUserNotFound
	}

	if user.Password != password {
		return "", ErrInvalidCredentials
	}

	body := tokenservice.TokenBody{
		Username:  user.Username,
		UserRole:  user.Role,
		SessionID: generateSessionID(user.ID),
	}

	token, err := tokenservice.GenerateToken(body, s.producer, s.expiry, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateSessionID creates a simple deterministic session identifier from the user ID.
// Replace with a proper UUID or random string generator in production.
func generateSessionID(userID int) string {
	return "session-" + itoa(userID)
}

// itoa is a minimal int-to-string helper to avoid importing strconv.
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	buf := make([]byte, 0, 10)
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		buf = append([]byte{byte('0' + i%10)}, buf...)
		i /= 10
	}
	if neg {
		buf = append([]byte{'-'}, buf...)
	}
	return string(buf)
}
