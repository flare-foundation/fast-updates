package provider_test

import (
	"fast-updates-client/provider"
	"testing"
)

func TestStringToDeltas(t *testing.T) {
	update := "+-0+"
	expected := []byte{113}
	result := provider.StringToDeltas(update)

	if len(result) != len(expected) {
		t.Errorf("Expected length of result to be %d, but got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] to be %d, but got %d", i, expected[i], result[i])
		}
	}
}
