package clc

import (
	"reflect"
	"testing"
)

func TestSearchName(t *testing.T) {
	if _, err := SearchName("AB0"); err == nil {
		t.Error("expected non-nil error; got nil")
	}

	result, err := SearchName("A666")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"A", "马克思列宁主义、毛泽东思想、邓小平理论"}) {
		t.Errorf("expected {A 马克思列宁主义、毛泽东思想、邓小平理论}; got %v", result)
	}

	result, err = SearchName("TP3")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"TP3", "计算技术、计算机技术"}) {
		t.Errorf("expected {TP3 计算技术、计算机技术}; got %v", result)
	}
}

func TestSearchCode(t *testing.T) {
	results := SearchCode("TP3")
	if l := len(results); l != 0 {
		t.Errorf("expected 0 results; got %d: %v", l, results)
	}

	results = SearchCode("计算机技术")
	if !reflect.DeepEqual(results, []CLC{{"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}}) {
		t.Errorf("expected [{TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}

func TestCLC(t *testing.T) {
	clc := CLC{"TP3", "计算技术、计算机技术"}
	if str := clc.String(); str != "TP3 计算技术、计算机技术" {
		t.Errorf("expected \"TP3 计算技术、计算机技术\"; got %q", str)
	}

	if top := clc.TopCategory(); !reflect.DeepEqual(top, CLC{"T", "工业技术"}) {
		t.Errorf("expected {T 工业技术}; got %v", top)
	}

	if results := clc.Categories(); !reflect.DeepEqual(
		results,
		[]CLC{{"T", "工业技术"}, {"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}},
	) {
		t.Errorf("expected [{T 工业技术} {TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}
