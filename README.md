Version
=======

This package helps stamp binary executables with version information.
When you build a go binary with this tool, you can print the version
information with the "version", "-version", or "--version" flag.

Usage
-----

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

    $ go get github.com/wx13/version

Build your code with:

    $ goversion build

The version number is obtained from a file called `.version` or `VERSION`
in the current directory or a parent directory.

Here is what the result looks like:

    $ ./foo version
    Version:    0.1
    Build Date: 2016-12-23T23:21:41-08:00
    Commit:     c4e91ed7dc70ec09f097b3847a88090e21dbe936

