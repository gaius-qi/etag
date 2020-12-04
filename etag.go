package etag

import (
	"crypto/sha1"
	"fmt"
)

// Generate the entity tag
func entitytag(entity string) string {
	// return quickly when empty
	if len(entity) == 0 {
		return "\"0-da39a3ee5e6b4b0d3255bfef95601890afd80709\""
	}

	// compute hash of entity
	hash := sha1.Sum([]byte(entity))

	return fmt.Sprintf("\"%d-%x\"", len(entity), hash)
}

// Generate a simple ETag
func Generate(entity string, weak bool) string {
	etag := entitytag(entity)

	if weak {
		return "W/" + etag
	}

	return etag
}
