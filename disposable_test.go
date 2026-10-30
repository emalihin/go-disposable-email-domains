package disposable

import "testing"

func TestIsDisposable(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{"user@0-mail.com", true},
		{"User@0-MAIL.COM", true},
		{"user@sub.0-mail.com", true},
		{"user@0-mail.com.", true},
		{"weird@local@0-mail.com", true},
		{"user@gmail.com", false},
		{"user@example.org", false},
		{"0-mail.com", false},
		{"user@", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := IsDisposable(tt.email); got != tt.want {
			t.Errorf("IsDisposable(%q) = %v, want %v", tt.email, got, tt.want)
		}
	}
}

func TestIsDisposableDomain(t *testing.T) {
	if !IsDisposableDomain("0-mail.com") {
		t.Error("expected 0-mail.com to be disposable")
	}
	if IsDisposableDomain("com") {
		t.Error("TLD alone must not be disposable")
	}
}

func TestListLoaded(t *testing.T) {
	if n := len(Domains()); n < 1000 {
		t.Fatalf("expected a substantial embedded list, got %d domains", n)
	}
}
