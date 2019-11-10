package propertyaccess

import (
	"testing"
)

func TestPropertyAccessor(t *testing.T) {
	propertyAccessor := NewPropertyAccessor()
	property := NewProperty("[customer][first_name]")

	value := "Marko's Rabbit Shop"

	customer := map[string]interface{}{
		"first_name": value,
	}

	var data = map[string]interface{}{
		"customer": customer,
	}

	extractedValue, err := propertyAccessor.GetValue(data, property)

	if err != nil {
		t.Errorf("GetValue failed with error: \"%s\"", err)
	}

	if value != extractedValue {
		t.Errorf("Expected value: \"%s\" doesn't match actual value: \"%s\"", value, extractedValue)
	}
}
