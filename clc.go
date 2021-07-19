package clc

import (
	"fmt"
	"strings"

	"github.com/sunshineplan/go-clc/class"
)

type CLC struct {
	code string
	name string
}

var TopCategory = []CLC{
	{"A", "马克思列宁主义、毛泽东思想、邓小平理论"},
	{"B", "哲学、宗教"},
	{"C", "社会科学总论"},
	{"D", "政治、法律"},
	{"E", "军事"},
	{"F", "经济"},
	{"G", "文化、科学、教育、体"},
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

func topCategory(code string) CLC {
	for _, i := range TopCategory {
		if i.code == code[:1] {
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

	code, ok := class.Verify(s[0])
	if !ok {
		panic(fmt.Sprintln("bad clc:", code))
	}

	return CLC{code, s[1]}
}

func (clc *CLC) Code() string {
	return clc.code
}

func (clc *CLC) Name() string {
	return clc.name
}

func (clc *CLC) String() string {
	return fmt.Sprintf("%s %s", clc.code, clc.name)
}

func (clc *CLC) TopCategory() CLC {
	return topCategory(clc.code)
}

func (clc *CLC) Categories() []CLC {
	categories := searchName(clc.code, class.LoadClass(clc.code))
	return append([]CLC{clc.TopCategory()}, categories...)
}

func searchName(code string, dict *[]class.Class) (results []CLC) {
	for _, i := range *dict {
		if strings.Contains(code, i.Code) {
			results = append(results, CLC{i.Code, i.Name})
			results = append(results, searchName(code, &i.SubClass)...)
		}
	}

	return
}

func SearchByCode(code string) (CLC, error) {
	var ok bool
	if code, ok = class.Verify(code); !ok {
		return CLC{}, fmt.Errorf("bad clc: %s", code)
	}

	results := searchName(code, class.LoadClass(code))

	if len(results) == 0 {
		return topCategory(code), nil
	}

	return results[len(results)-1], nil
}

func SearchByName(name string) (result []CLC) {
	for _, i := range class.MatchAll(name) {
		if clc := str2clc(i); strings.Contains(clc.name, name) {
			result = append(result, clc)
		}
	}

	return
}
