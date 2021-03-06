package main

import (
	"fmt"
	"github.com/wx13/version.v0"
)

func main() {
	// Use a custom print template.
	p := version.NewPrinter()
	p.FullTemplate = "This is version {{.Version}}, " +
		"built at {{.Date}}, from commit {{.Commit}}."
	p.Template = "version: {{.Version}}"
	p.Print()

	// Use version information elsewhere in the app
	fmt.Printf("[%s] Do something...\n", version.Version)
	fmt.Printf("[%s] Do something else...\n", version.FullVersion)
}
