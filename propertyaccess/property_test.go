package propertyaccess

import (
	"testing"
)

func TestProperty(t *testing.T) {
	property := NewProperty("[customer][first_name]")

	length := len(property.Path())

	if 2 != length {
		t.Errorf("Expected property length was 2, got \"%d\" instead.", length)
	}
}
