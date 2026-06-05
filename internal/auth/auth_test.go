package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-super-secret-key"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	extractedID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if extractedID != userID {
		t.Errorf("expected user ID %v, got %v", userID, extractedID)
	}
}

func TestValidateExpiredJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-super-secret-key"
	// Set a very short expiration time
	expiresIn := time.Millisecond

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	// Sleep to let the token expire
	time.Sleep(10 * time.Millisecond)

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatal("expected error validating expired token, got nil")
	}
}

func TestValidateWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "my-super-secret-key"
	wrongSecret := "wrong-secret"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	// Attempt to validate with the wrong secret key
	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Fatal("expected error validating token with wrong secret, got nil")
	}
}

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-super-secret-token")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token != "my-super-secret-token" {
		t.Errorf("expected my-super-secret-token, got %s", token)
	}

	// Test missing header
	emptyHeaders := http.Header{}
	_, err = GetBearerToken(emptyHeaders)
	if err == nil {
		t.Error("expected error for missing header, got nil")
	}

	// Test malformed header
	badHeaders := http.Header{}
	badHeaders.Set("Authorization", "Basic my-super-secret-token")
	_, err = GetBearerToken(badHeaders)
	if err == nil {
		t.Error("expected error for malformed header, got nil")
	}
}
