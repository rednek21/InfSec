package stribog

import (
	"encoding/hex"
	"testing"
)

func TestStribog512_SameInputSameHash(t *testing.T) {
	message := "test input"

	h1 := New(64)
	h1.Write([]byte(message))
	hash1 := h1.Sum(nil)

	h2 := New(64)
	h2.Write([]byte(message))
	hash2 := h2.Sum(nil)

	if hex.EncodeToString(hash1) != hex.EncodeToString(hash2) {
		t.Errorf("Expected hashes to be the same, but got different hashes: %s vs %s", hex.EncodeToString(hash1), hex.EncodeToString(hash2))
	}
}
