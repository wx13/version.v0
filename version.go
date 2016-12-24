package version

import (
	"os"
)

var (
	Version string
	Date    string
	Commit  string
)

func PrintVersion() {
	println(Version)
}

func Run() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "version", "-version", "--version":
		PrintVersion()
		os.Exit(0)
	}
}
