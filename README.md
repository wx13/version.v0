Version
=======

This package helps stamp binary executables with version information.
When you build a go binary with this tool, you can print the version
information with the "version", "-version", or "--version" flag.

Usage
-----

First get the package:

    $ go get github.com/wx13/version

Then add the line

    version.Run()

to your main function or an init function.  Finaly, build your code with:

    $ goversion build

Any arguments after "goversion build" are passed to the go tool.  Now you
can do:

    $ ./foo version
    Version:    0.1
    Build Date: 2016-12-23T23:21:41-08:00
    Commit:     c4e91ed7dc70ec09f097b3847a88090e21dbe936

The version number is obtained from a file called `.version` in the current
directory or a parent directory.

Example
-------

See the `example` directory for a working example.  Here is what a minimal
program looks like:

    package main

    import (
        "github.com/wx13/version"
    )

    func main() {
        version.Run()
        prntln("Do something...")
    }
