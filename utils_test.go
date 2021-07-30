package flag

import (
	"github.com/sleeyax/flagstruct/internal/tests"
	"testing"
)

func TestToField(t *testing.T) {
	flag := "this-is-a-field"
	expected := "ThisIsAField"

	if actual := ToField(flag); actual != expected {
		t.Fatalf(tests.MismatchFormat, expected, actual)
	}
}

func TestToFlag(t *testing.T) {
	field := "ThisIsAField"
	expected := "this-is-a-field"

	if actual := ToFlag(field); actual != expected {
		t.Fatalf(tests.MismatchFormat, expected, actual)
	}
}
