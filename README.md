# go-clc

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]
[![Go](https://github.com/sunshineplan/go-clc/workflows/Test/badge.svg)][actions]

[godev]: https://pkg.go.dev/github.com/sunshineplan/go-clc "Documentation"
[actions]: https://github.com/sunshineplan/go-clc/actions "GitHub Actions Page"

Golang Chinese Library Classification (CLC; 中国图书馆分类法) Search Tool

All classification data were collected from [wikipedia](https://zh.wikipedia.org/wiki/%E4%B8%AD%E5%9B%BD%E5%9B%BE%E4%B9%A6%E9%A6%86%E5%88%86%E7%B1%BB%E6%B3%95).

《中国图书馆分类法》各种版本及其管理系统的知识产权归国家图书馆所有。http://clc.nlc.cn

## Installation

    go get -u github.com/sunshineplan/go-clc

## License

[The MIT License (MIT)](https://raw.githubusercontent.com/sunshineplan/go-clc/main/LICENSE)

## Example code

```go
package main

import (
	"log"

	"github.com/sunshineplan/go-clc"
)

func main() {
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
```
