package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	id := uuid.New()
	tok, err := MakeJWT(id, "secret", time.Hour)

	if err != nil || tok == "" {
		t.Fatalf("expected token, got %q, err%v", tok, err)
	}
}

func TestValidJWT(t *testing.T) {
	id := uuid.New()
	secret := "secret"

	good, _ := MakeJWT(id, secret, time.Hour)
	expired, _ := MakeJWT(id, secret, -time.Minute)

	tests := []struct {
		name    string
		token   string
		secret  string
		wantErr bool
	}{
		{"valid", good, secret, false},
		{"expired", expired, secret, true},
		{"wrong secret", good, "other", true},
		{"garbage", "not.a.token", secret, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ValidateJWT(tc.token, tc.secret)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr && got != id {
				t.Fatalf("got %v, want %v", got, id)
			}
		})
	}
}
