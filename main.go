package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/grake/cmd/clean"
	"github.com/TheDevtop/grake/cmd/setup"
)

func argShift() {
	os.Args = os.Args[1:]
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: grake [command] [flags]\nCommands: init|build|clean")
	os.Exit(2)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	argShift()
	switch os.Args[0] {
	case "init":
		setup.CmdMain()
	case "build":
		setup.CmdMain()
	case "clean":
		clean.CmdMain()
	default:
		usage()
	}
}
