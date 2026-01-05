package mpr

import (
	"crypto/rand"
	"fmt"
)

func init() {
	// Use secure random UUID generation
	randomRead = func(b []byte) (int, error) {
		return rand.Read(b)
	}
}

// GenerateID generates a new unique ID for model elements.
func GenerateID() string {
	return generateUUID()
}

// Hash computes a hash for content (used for content deduplication).
func Hash(content []byte) string {
	// Simple hash for now - could use crypto/sha256 for better hashing
	var sum uint64
	for i, b := range content {
		sum += uint64(b) * uint64(i+1)
	}
	return fmt.Sprintf("%016x", sum)
}

// ValidateID checks if an ID is valid.
func ValidateID(id string) bool {
	if len(id) != 36 {
		return false
	}
	// Check UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	for i, c := range id {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if c != '-' {
				return false
			}
		} else {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return false
			}
		}
	}
	return true
}
