package clc_test

import (
	"log"

	"github.com/sunshineplan/go-clc"
)

func Example() {
	results := clc.SearchByName("计算机技术")
	log.Print(results) // [{TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]

	result, err := clc.SearchByCode("TP3")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)               // {TP3 计算技术、计算机技术}
	log.Print(result.String())      // TP3 计算技术、计算机技术
	log.Print(result.TopCategory()) // {T 工业技术}
	log.Print(result.Categories())  // [{T 工业技术} {TP 自动化技术、计算机技术} {TP3 计算技术、计算机技术}]
}
