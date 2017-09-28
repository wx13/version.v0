package main

import (
	"fmt"
	"github.com/wx13/version"
)

func main() {
	// Use a custom print template.
	p := version.NewPrinter()
	p.Template = "This is version {{.Version}}, " +
		"built at {{.Date}}, from commit {{.Commit}}."
	p.Print()

	// Use version information elsewhere in the app
	fmt.Printf("[%s] Do something...\n", version.Version)
}
