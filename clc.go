package clc

import (
	"fmt"
	"strings"

	"github.com/sunshineplan/go-clc/class"
)

// CLC represents Chinese Library Classification structure.
type CLC struct {
	// A notation is a code commonly used in classification schemes to represent a class.
	notation string
	// class description.
	caption string
}

// TopClass is a list of top-level class defined in Chinese Library Classification.
var TopClass = []CLC{
	{"A", "马克思列宁主义、毛泽东思想、邓小平理论"},
	{"B", "哲学、宗教"},
	{"C", "社会科学总论"},
	{"D", "政治、法律"},
	{"E", "军事"},
	{"F", "经济"},
	{"G", "文化、科学、教育、体育"},
	{"H", "语言、文字"},
	{"I", "文学"},
	{"J", "艺术"},
	{"K", "历史、地理"},
	{"N", "自然科学总论"},
	{"O", "数理科学和化学"},
	{"P", "天文学、地球科学"},
	{"Q", "生物科学"},
	{"R", "医药、卫生"},
	{"S", "农业科学"},
	{"T", "工业技术"},
	{"U", "交通运输"},
	{"V", "航空、航天"},
	{"X", "环境科学、安全科学"},
	{"Z", "综合性图书"},
}

func topClass(notation string) CLC {
	for _, i := range TopClass {
		if i.notation == notation[:1] {
			return i
		}
	}

	return CLC{}
}

func str2clc(str string) CLC {
	s := strings.SplitN(str, " ", 2)
	if len(s) != 2 {
		panic(fmt.Sprintln("failed to convert string to clc:", str))
	}

	notation, ok := class.Verify(s[0])
	if !ok {
		panic(fmt.Sprintln("bad CLC notation:", notation))
	}

	return CLC{notation, s[1]}
}

// Notation returns CLC's notation.
func (clc *CLC) Notation() string {
	return clc.notation
}

// Caption returns CLC's caption.
func (clc *CLC) Caption() string {
	return clc.caption
}

// String returns CLC's notation and caption in string format.
func (clc *CLC) String() string {
	return fmt.Sprintf("%s %s", clc.notation, clc.caption)
}

// TopClass returns CLC's top-level class.
func (clc *CLC) TopClass() CLC {
	return topClass(clc.notation)
}

// Classes returns CLC's all related classes.
func (clc *CLC) Classes() []CLC {
	classes := searchCaption(clc.notation, class.LoadClass(clc.notation))
	return append([]CLC{clc.TopClass()}, classes...)
}

func searchCaption(notation string, dict *[]class.Class) (results []CLC) {
	for _, i := range *dict {
		if strings.Contains(notation, i.Notation) {
			results = append(results, CLC{i.Notation, i.Caption})
			results = append(results, searchCaption(notation, &i.SubClass)...)
		}
	}

	return
}

// SearchByNotation searchs CLC by notation and returns most likely result.
// The only possible returned error when notation is illegal.
func SearchByNotation(notation string) (CLC, error) {
	var ok bool
	if notation, ok = class.Verify(notation); !ok {
		return CLC{}, fmt.Errorf("bad CLC notation: %s", notation)
	}

	results := searchCaption(notation, class.LoadClass(notation))

	if len(results) == 0 {
		return topClass(notation), nil
	}

	return results[len(results)-1], nil
}

// SearchByCaption searchs CLC by caption and returns a slice of all
// successive matches. A return value of nil indicates no match or caption
// is empty string.
func SearchByCaption(caption string) (results []CLC) {
	if len(caption) == 0 {
		return
	}

	for _, i := range class.FindAll(caption) {
		if clc := str2clc(i); strings.Contains(clc.caption, caption) {
			results = append(results, clc)
		}
	}

	return
}
