package main

import (
	"flag"

	"github.com/sunshineplan/go-clc/class"
)

var dir string
var debug bool

func main() {
	flag.StringVar(&dir, "dir", "data", "directory to export")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	class.ExportJSON(dir, debug)
}
