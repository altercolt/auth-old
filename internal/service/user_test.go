package service

import (
	"testing"
)

func TestSalt(t *testing.T) {
	salt := generateSalt()
	t.Logf("salt : %s", salt)
}
