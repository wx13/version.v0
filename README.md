# Version

This package helps stamp binary executables with version information. When you
build a go binary with this tool, you can print the version information with the
"version", "-version", or "--version" flag.

## Usage

Here is what a minimal program looks like:

    package main

    import (
        "github.com/wx13/version"
    )

    func main() {
        version.Run()
        println("Do something...")
    }

Now go get the package:

    $ go get github.com/wx13/version/goversion

Build your code with:

    $ goversion build <any other normal build args>

The version number is obtained from a file called `.version` or `VERSION` in the
current directory or a parent directory.

Here is what the result looks like:

    $ ./foo version
    Version:    0.1
    Build Date: 2016-12-23T23:21:41-08:00
    Commit:     c4e91ed7dc70ec09f097b3847a88090e21dbe936

## Advanced Usage

See the example directory for more advanced usage. You can set a custom print
template. You can also use version information in your program, such as in a UI
widget.
