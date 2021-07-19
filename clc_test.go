package clc

import (
	"reflect"
	"testing"
)

func TestSearchByNotation(t *testing.T) {
	if _, err := SearchByNotation("AB0"); err == nil {
		t.Error("expected non-nil error; got nil")
	}

	result, err := SearchByNotation("A666")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"A", "马克思列宁主义、毛泽东思想、邓小平理论"}) {
		t.Errorf("expected {A 马克思列宁主义、毛泽东思想、邓小平理论}; got %v", result)
	}

	result, err = SearchByNotation("TP3")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"TP3", "计算技术、计算机技术"}) {
		t.Errorf("expected {TP3 计算技术、计算机技术}; got %v", result)
	}
}

func TestSearchByCaption(t *testing.T) {
	results := SearchByCaption("TP3")
	if l := len(results); l != 0 {
		t.Errorf("expected 0 results; got %d: %v", l, results)
	}

	results = SearchByCaption("计算机技术")
	if !reflect.DeepEqual(results, []CLC{{"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}}) {
		t.Errorf("expected [{TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}

func TestCLC(t *testing.T) {
	clc := CLC{"TP3", "计算技术、计算机技术"}
	if str := clc.String(); str != "TP3 计算技术、计算机技术" {
		t.Errorf("expected \"TP3 计算技术、计算机技术\"; got %q", str)
	}

	if class := clc.TopClass(); !reflect.DeepEqual(class, CLC{"T", "工业技术"}) {
		t.Errorf("expected {T 工业技术}; got %v", class)
	}

	if results := clc.Classes(); !reflect.DeepEqual(
		results,
		[]CLC{{"T", "工业技术"}, {"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}},
	) {
		t.Errorf("expected [{T 工业技术} {TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}
