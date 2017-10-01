# Version

This package helps stamp binary executables with version information. When you
build a go binary with this tool, you can print the version information with the
"version", "-version", or "--version" flag.

## Motivation

When I build a binary, I want the binary to report its version (on the
command-line or in the gui). I want release versions to report a version
like 'v0.5.2', but development versions should provide more information
(build time stamp, commit hash) to aid in debugging.

## How it works.

There are two parts to 'version'. First there is a library which defines
variables (version, build date, etc) and some convenience functions for
printing version info. Then there is a binary (`goversion`) which wraps
the `go` binary and inserts version info using ldflags.

The `goversion` binary gets the version in the following way. It looks
for git tags that start with 'v' followed by a digit. Then it finds the
most recent tag which is an ancestor of HEAD. If it doesn't find such a
tag, it looks for a file called `.version` or `VERSION` in the current
directory or any parent directory.

Next it looks at the age of the version tag. If it's current then it uses
just the tag name as the version. If it's not current, then it adds a
partial git commit hash to the version string (e.g. 'v0.3.2-a83bc229').

## Usage

Here is what a minimal program looks like:

    package main

    import "github.com/wx13/version"

    func main() {
        version.Print()
        println("Do something...")
    }

Now go get the package:

    $ go get github.com/wx13/version.v0/goversion

Build your code with:

    $ goversion build <any other normal build args>

Here is what the result might look like:

    $ ./foo version
    Version:    0.1
    Build Date: 2016-12-23T23:21:41-08:00
    Commit:     c4e91ed7dc70ec09f097b3847a88090e21dbe936


## Advanced Usage

See the example directory for more advanced usage. You can set a custom print
template. You can also use version information in your program, such as in a UI
widget.
