package tokenservice

import (
	"testing"
	"time"
)

const testSecret = "super-secret-key"

func TestGenerateAndValidateToken(t *testing.T) {
	body := TokenBody{
		Username:  "akif",
		UserRole:  "admin",
		SessionID: "session-123",
	}

	token, err := GenerateToken(body, "game-server", time.Minute*5, testSecret)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	parsed, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("failed to validate token: %v", err)
	}

	if parsed.Body.Username != body.Username {
		t.Errorf("expected username %s, got %s", body.Username, parsed.Body.Username)
	}
}

func TestInvalidSignature(t *testing.T) {
	body := TokenBody{
		Username:  "akif",
		UserRole:  "admin",
		SessionID: "session-123",
	}

	token, err := GenerateToken(body, "game-server", time.Minute*5, testSecret)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateToken(token, "wrong-secret")
	if err == nil {
		t.Fatal("expected invalid signature error")
	}
}

func TestExpiredToken(t *testing.T) {
	body := TokenBody{
		Username:  "akif",
		UserRole:  "admin",
		SessionID: "session-123",
	}

	token, err := GenerateToken(body, "game-server", -time.Minute, testSecret)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateToken(token, testSecret)
	if err == nil {
		t.Fatal("expected expiration error")
	}
}

// TestRoleInToken verifies that the UserRole claim survives round-trip serialization.
func TestRoleInToken(t *testing.T) {
	expectedRole := "moderator"
	body := TokenBody{
		Username:  "alice",
		UserRole:  expectedRole,
		SessionID: "session-abc",
	}

	token, err := GenerateToken(body, "game-server", time.Minute*5, testSecret)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	parsed, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("failed to validate token: %v", err)
	}

	if parsed.Body.UserRole != expectedRole {
		t.Errorf("expected role %q, got %q", expectedRole, parsed.Body.UserRole)
	}
}

// TestSessionIDInToken verifies that the SessionID claim survives round-trip serialization.
func TestSessionIDInToken(t *testing.T) {
	expectedSessionID := "session-xyz-789"
	body := TokenBody{
		Username:  "bob",
		UserRole:  "player",
		SessionID: expectedSessionID,
	}

	token, err := GenerateToken(body, "game-server", time.Minute*5, testSecret)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	parsed, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("failed to validate token: %v", err)
	}

	if parsed.Body.SessionID != expectedSessionID {
		t.Errorf("expected session ID %q, got %q", expectedSessionID, parsed.Body.SessionID)
	}
}
