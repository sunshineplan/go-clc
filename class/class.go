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

var re = regexp.MustCompile(fmt.Sprintf(`^%s%s$`, prefix, suffix))

// Verify verifies notation is legal or not.
func Verify(notation string) (string, bool) {
	notation = strings.ToUpper(notation)

	return notation, re.MatchString(notation)
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

func keys() (keys []string) {
	for k := range index {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return
}

type class string

// Class represents Chinese Library Classification structure with subclass.
type Class struct {
	Notation string
	Caption  string
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

		notation, ok := Verify(s[0])
		if !ok {
			panic(fmt.Sprintln("bad CLC notation:", notation))
		}

		if dst := findClass(&class, notation); dst != nil {
			if debug {
				log.Println("found:", line, dst.Notation)
			}
			(*dst).SubClass = append((*dst).SubClass, Class{
				Notation: notation,
				Caption:  s[1],
			})
		} else {
			if debug {
				log.Println("not found:", line)
			}
			class = append(class, Class{
				Notation: notation,
				Caption:  s[1],
			})
		}
	}

	return class
}

// Used to store loaded class data.
var cache sync.Map

// LoadClass loads class data according str's first letter.
// The loaded data will be stored in cache, so the same class
// data can be loaded faster next time.
func LoadClass(str string) *[]Class {
	if len(s) == 0 {
		panic("argument must not be empty string")
	}

	notation, ok := Verify(str[:1])
	if !ok {
		panic(fmt.Sprintln("bad CLC notation:", notation))
	}

	value, ok := cache.Load(notation)
	if !ok {
		class := index[notation]
		data := class.load(false)

		cache.Store(notation, &data)

		return &data
	}

	return value.(*[]Class)
}

func findClass(src *[]Class, notation string) *Class {
	if len(notation) == 0 {
		return nil
	}

	var found *Class
	for i := range *src {
		if reduceOne(notation) == (*src)[i].Notation {
			return &(*src)[i]
		} else if strings.Contains(reduceOne(notation), (*src)[i].Notation) {
			found = &(*src)[i]
			if dst := findClass(&(*src)[i].SubClass, notation); dst != nil {
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

// FindAll walks all classes raw data and returns a slice of all
// successive matches of str. A return value of nil indicates no match.
func FindAll(str string) (results []string) {
	for _, i := range keys() {
		for _, line := range strings.Split(string(index[i]), "\n") {
			if strings.Contains(line, str) {
				results = append(results, line)
			}
		}
	}

	return
}

func exportJSON(dir, class string, debug bool) {
	path := fmt.Sprintf("%s/%s.json", dir, class)
	log.Println("Exporting", path)

	data := index[class].load(debug)

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

// ExportJSON exports all classes data to separate files in json format.
func ExportJSON(dir string, debug bool) {
	for _, class := range keys() {
		exportJSON(dir, class, debug)
	}
}
