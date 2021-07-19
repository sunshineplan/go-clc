package main

import (
	"flag"

	"github.com/sunshineplan/go-clc/class"
)

var dir string
var force, debug bool

func main() {
	flag.StringVar(&dir, "dir", "data", "directory to export")
	flag.BoolVar(&force, "force", false, "force overwrite existing files")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	class.ExportJSON(dir, force, debug)
}
