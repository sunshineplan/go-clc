package class

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	class := findIndex(&[]Class{
		{Code: "TQ"},
	}, "TQ17")
	if class == nil {
		t.Fatal("expected non-nil; got nil")
	}
}
