package valueobjects

import "testing"

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		wantErr   error
	}{
		{name: "valid password", plaintext: "supersecret", wantErr: nil},
		{name: "empty password", plaintext: "", wantErr: ErrPasswordEmpty},
		{name: "too short", plaintext: "short", wantErr: ErrPasswordTooShort},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd, err := NewPassword(tt.plaintext)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				if pwd != nil {
					t.Fatalf("expected nil password on error, got %v", pwd)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if pwd.Hash() == tt.plaintext {
				t.Fatal("password hash must not equal the plaintext")
			}
			if err := pwd.Verify(tt.plaintext); err != nil {
				t.Fatalf("expected verify to succeed: %v", err)
			}
		})
	}
}

func TestPassword_Verify(t *testing.T) {
	pwd, err := NewPassword("correct-horse-battery")
	if err != nil {
		t.Fatalf("unexpected error creating password: %v", err)
	}

	if err := pwd.Verify("correct-horse-battery"); err != nil {
		t.Fatalf("expected verify to succeed for correct password: %v", err)
	}

	if err := pwd.Verify("wrong-password"); err == nil {
		t.Fatal("expected verify to fail for incorrect password")
	}
}

func TestNewPasswordFromHash(t *testing.T) {
	valid, err := NewPassword("hunter22")
	if err != nil {
		t.Fatalf("unexpected error creating password: %v", err)
	}

	tests := []struct {
		name    string
		hash    string
		wantErr error
	}{
		{name: "empty hash", hash: "", wantErr: ErrInvalidHash},
		{name: "malformed hash", hash: "not-a-bcrypt-hash", wantErr: ErrInvalidHash},
		{name: "valid hash", hash: valid.Hash(), wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd, err := NewPasswordFromHash(tt.hash)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if err := pwd.Verify("hunter22"); err != nil {
				t.Fatalf("expected reconstructed password to verify: %v", err)
			}
		})
	}
}

func TestPassword_Matches(t *testing.T) {
	a, _ := NewPassword("password-one")
	b, _ := NewPasswordFromHash(a.Hash())
	c, _ := NewPassword("password-two")

	if !a.Matches(b) {
		t.Fatal("expected passwords with identical hash to match")
	}
	if a.Matches(c) {
		t.Fatal("expected passwords with different hashes to not match")
	}
}

func TestPassword_MarshalJSON_DoesNotLeakHash(t *testing.T) {
	pwd, _ := NewPassword("do-not-leak-me")

	data, err := pwd.MarshalJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(data) != `"Password(***)"`+"\n" {
		t.Fatalf("expected redacted marshaled output, got %q", data)
	}
}

func TestValidatePassword(t *testing.T) {
	if err := ValidatePassword(""); err != ErrPasswordEmpty {
		t.Fatalf("expected ErrPasswordEmpty, got %v", err)
	}
	if err := ValidatePassword("short"); err != ErrPasswordTooShort {
		t.Fatalf("expected ErrPasswordTooShort, got %v", err)
	}
	if err := ValidatePassword("longenough"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
