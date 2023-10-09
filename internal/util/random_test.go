//go:build !integration

package util

import (
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	max := int64(10)
	val := RandomInt(max)
	if val < 0 || val >= max {
		t.Errorf("RandomInt out of range: got %v, expected between 0 and %v", val, max)
	}
}

func TestRandomString(t *testing.T) {
	length := 5
	str := RandomString(length)
	if len(str) != length {
		t.Errorf("RandomString length mismatch: got %v, expected %v", len(str), length)
	}

	for _, char := range str {
		if !strings.Contains(
			alphabet,
			string(char),
		) {
			t.Errorf("RandomString contains invalid character: got %v, expected characters from %v", char, alphabet)
		}
	}
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	if len(owner) != 6 {
		t.Errorf("RandomOwner length mismatch: got %v, expected 6", len(owner))
	}
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()
	if money < 0 || money >= 1000 {
		t.Errorf("RandomMoney out of range: got %v, expected between 0 and 1000", money)
	}
}

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		t.Errorf("RandomEmail format invalid: got %v, expected [string]@email.com", email)
	}

	if len(parts[0]) != 6 {
		t.Errorf("RandomEmail local part length mismatch: got %v, expected 6", len(parts[0]))
	}

	if parts[1] != "email.com" {
		t.Errorf("RandomEmail domain mismatch: got %v, expected email.com", parts[1])
	}
}
