package clc

import (
	"reflect"
	"testing"
)

func TestSearchByCode(t *testing.T) {
	if _, err := SearchByCode("AB0"); err == nil {
		t.Error("expected non-nil error; got nil")
	}

	result, err := SearchByCode("A666")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"A", "马克思列宁主义、毛泽东思想、邓小平理论"}) {
		t.Errorf("expected {A 马克思列宁主义、毛泽东思想、邓小平理论}; got %v", result)
	}

	result, err = SearchByCode("TP3")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, CLC{"TP3", "计算技术、计算机技术"}) {
		t.Errorf("expected {TP3 计算技术、计算机技术}; got %v", result)
	}
}

func TestSearchByName(t *testing.T) {
	results := SearchByName("TP3")
	if l := len(results); l != 0 {
		t.Errorf("expected 0 results; got %d: %v", l, results)
	}

	results = SearchByName("计算机技术")
	if !reflect.DeepEqual(results, []CLC{{"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}}) {
		t.Errorf("expected [{TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}

func TestCLC(t *testing.T) {
	clc := CLC{"TP3", "计算技术、计算机技术"}
	if str := clc.String(); str != "TP3 计算技术、计算机技术" {
		t.Errorf("expected \"TP3 计算技术、计算机技术\"; got %q", str)
	}

	if category := clc.TopCategory(); !reflect.DeepEqual(category, CLC{"T", "工业技术"}) {
		t.Errorf("expected {T 工业技术}; got %v", category)
	}

	if results := clc.Categories(); !reflect.DeepEqual(
		results,
		[]CLC{{"T", "工业技术"}, {"TP", "自动化技术、计算机技术"}, {"TP3", "计算技术、计算机技术"}},
	) {
		t.Errorf("expected [{T 工业技术} {TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]; got %v", results)
	}
}
