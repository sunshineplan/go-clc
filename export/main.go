package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/sunshineplan/go-clc/class"
)

var (
	dir   = flag.String("dir", "data", "directory to export")
	debug = flag.Bool("debug", false, "debug")
)

func main() {
	flag.Parse()
	if *debug {
		level := new(slog.LevelVar)
		level.Set(slog.LevelDebug)
		slog.SetDefault(slog.New(slog.NewTextHandler(log.Default().Writer(), &slog.HandlerOptions{Level: level})))
	}

	class.ExportJSON(*dir)
}
