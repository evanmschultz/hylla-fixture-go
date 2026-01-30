package beta

import "github.com/google/uuid"

// NewID returns a new UUID string for fixture identifiers.
func NewID() string {
	return uuid.NewString()
}
