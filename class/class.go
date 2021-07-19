package class

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

const (
	prefix = "([A-K]|[N-S]|T(B|[D-H]|[J-N]|P|Q|S|U|V)?|U|V|X|Z)"
	suffix = `(\d{0,3}(\.\d)?-)?\d{0,3}(\+\d{0,1})?(\.\d{1,3}(\+\d{0,2})?)?(\.\d{1,2})?(/-?\d{0,3}\.?\d{1,3})?`
)

var reCLC = regexp.MustCompile(fmt.Sprintf(`^%s%s$`, prefix, suffix))

func Verify(code string) (string, bool) {
	code = strings.ToUpper(code)

	return code, reCLC.MatchString(code)
}

var index = map[string]class{
	"A": a,
	"B": b,
	"C": c,
	"D": d,
	"E": e,
	"F": f,
	"G": g,
	"H": h,
	"I": i,
	"J": j,
	"K": k,
	"N": n,
	"O": o,
	"P": p,
	"Q": q,
	"R": r,
	"S": s,
	"T": t,
	"U": u,
	"V": v,
	"X": x,
	"Z": z,
}

type class string

type Class struct {
	Code     string
	Name     string
	SubClass []Class `json:",omitempty"`
}

func (c class) load(debug bool) []Class {
	src := strings.Split(strings.TrimSpace(string(c)), "\n")
	sort.Strings(src)

	var class []Class
	for _, line := range src {
		s := strings.SplitN(line, " ", 2)
		if len(s) != 2 {
			panic(fmt.Sprintln("failed to convert string:", line))
		}

		code, ok := Verify(s[0])
		if !ok {
			panic(fmt.Sprintln("bad clc:", code))
		}

		if dst := findIndex(&class, code); dst != nil {
			if debug {
				log.Println("found:", line, dst.Code)
			}
			(*dst).SubClass = append((*dst).SubClass, Class{
				Code: code,
				Name: s[1],
			})
		} else {
			if debug {
				log.Println("not found:", line)
			}
			class = append(class, Class{
				Code: code,
				Name: s[1],
			})
		}
	}

	return class
}

var cache sync.Map

func LoadClass(s string) *[]Class {
	key := s[:1]
	value, ok := cache.Load(key)
	if !ok {
		class := index[key]
		data := class.load(false)

		cache.Store(key, &data)

		return &data
	}

	return value.(*[]Class)
}

func findIndex(src *[]Class, code string) *Class {
	if len(code) == 0 {
		return nil
	}

	var found *Class
	for i := range *src {
		if reduceOne(code) == (*src)[i].Code {
			return &(*src)[i]
		} else if strings.Contains(reduceOne(code), (*src)[i].Code) {
			found = &(*src)[i]
			if dst := findIndex(&(*src)[i].SubClass, code); dst != nil {
				return dst
			}
		}
	}

	return found
}

func reduceOne(s string) string {
	if strings.Contains(s, "/") {
		return s[:strings.Index(s, "/")-1]
	}

	return regexp.MustCompile(`(\.|-|\+)+$`).ReplaceAllString(s[:len(s)-1], "")
}

func MatchAll(str string) (results []string) {
	for _, i := range index {
		for _, line := range strings.Split(string(i), "\n") {
			if strings.Contains(line, str) {
				results = append(results, line)
			}
		}
	}

	return
}

func exportJSON(dir, category string, debug bool) {
	path := fmt.Sprintf("%s/%s.json", dir, category)
	log.Println("Exporting", path)

	class := index[category]
	data := class.load(debug)

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}
	dst, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	if _, err := dst.Write(b); err != nil {
		log.Fatal(err)
	}
}

func ExportJSON(dir string, force, debug bool) {
	var keys []string
	for k := range index {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, category := range keys {
		if !force {
			path := fmt.Sprintf("%s/%s.json", dir, category)
			if _, err := os.Stat(path); err == nil {
				log.Println("Skip export", path)
				continue
			}
		}
		exportJSON(dir, category, debug)
	}
}
