// Package version provides a simple way to version an executable.
package version

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Use package variables because they can be set linker flags.
var (
	Version     string
	FullVersion string
	Date        string
	Commit      string
)

// Printer allows the user to set a custom print template.
type Printer struct {
	Template string
	Flags    []string
}

// NewPrinter sets the default print template.
func NewPrinter() *Printer {
	p := Printer{}
	p.Flags = []string{"version", "-version", "--version"}
	p.Template = ""
	p.Template += "Version:    {{.Version}}\n"
	p.Template += "Build Date: {{.Date}}\n"
	p.Template += "Commit:     {{.Commit}}\n"
	return &p
}

// FlagIsSet returns true if a version flag is set.
func (p *Printer) FlagIsSet(arg string) bool {
	for _, flag := range p.Flags {
		if arg == flag {
			return true
		}
	}
	return false
}

// Print prints version info, conditioned on the user input.
func (p *Printer) Print() error {
	// Don't do anything if no arguments are passed.
	if len(os.Args) < 2 {
		return nil
	}
	// Check if user is asking for version information.
	if p.FlagIsSet(os.Args[1]) {

		// Compile the print template.
		t := template.Must(template.New("VersionPrinter").Parse(p.Template))
		var b bytes.Buffer

		// Execute the template and output the result.
		err := t.Execute(&b, map[string]string{
			"Version": Version,
			"Date":    Date,
			"Commit":  Commit,
		})
		if err != nil {
			return err
		}
		fmt.Println(b.String())

		os.Exit(0)
	}
	return nil
}

// Print is a shortcut to printing with the default template.
func Print() {
	p := NewPrinter()
	p.Print()
}
