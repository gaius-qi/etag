package etag

import (
	"testing"
)

func TestGenerateEtag(t *testing.T) {
	entitys := []string{"etag", ""}
	etags := []string{
		"\"4-92e200311a56800b3e475bf2d2442724535e87bf\"",
		"\"6-ae0a7afecee1e83d27e322e287cb9bf351ea0018\"",
		"\"0-da39a3ee5e6b4b0d3255bfef95601890afd80709\"",
	}

	for i, entity := range entitys {
		etag := Generate(entity, false)
		if etag != etags[i] {
			t.Errorf("%s: %s != %s", entity, etag, etags[i])
		}
	}
}

func TestGenerateWeekEtag(t *testing.T) {
	entitys := []string{"etag", ""}
	etags := []string{
		"W/\"4-92e200311a56800b3e475bf2d2442724535e87bf\"",
		"W/\"6-ae0a7afecee1e83d27e322e287cb9bf351ea0018\"",
		"W/\"0-da39a3ee5e6b4b0d3255bfef95601890afd80709\"",
	}

	for i, entity := range entitys {
		etag := Generate(entity, true)
		if etag != etags[i] {
			t.Errorf("%s: %s != %s", entity, etag, etags[i])
		}
	}
}
