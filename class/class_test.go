package class

import (
	"testing"
)

func TestFindClass(t *testing.T) {
	class := findClass(&[]Class{
		{Notation: "TQ"},
	}, "TQ17")
	if class == nil {
		t.Fatal("expected non-nil; got nil")
	}
}
