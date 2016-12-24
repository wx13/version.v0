package version

import (
	"fmt"
	"os"
)

var (
	Version string
	Date    string
	Commit  string
)

// PrintVersion prints version info to stdout.
func PrintVersion() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Build Date: %s\n", Date)
	fmt.Printf("Commit:     %s\n", Commit)
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
