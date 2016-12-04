Go 101
======

History
-------

- Designed at Google 7 years ago
- By Robert Griesemer, Rob Pike & Ken Thompson
- To improve (and because they hated) C++

Why YES?
--------

- Async concurrency as part of the language
- Compiled to a statically linked binary
- Statically typed
- Garbage collected
- Fast
- Go tools!

Why NO?
-------

You probably heard about:

- No generics
- Dependency managment
- Libraries not always as mature, but improving!

GOPATH
------

Your project will usually have the following struct:

    $GOPATH/src/github.com/user/project
    
Example:

    ~/go/src/github.com/agonzalezro/polo

Vendoring
---------

- Since Go 1.5: `$GO15VENDOREXPERIMENT`

- Utils to manage vendor:

  + [`godeps`](https://github.com/tools/godep) you track deps in your VCS: `{get,save,restore}`
  + [`glide`](https://github.com/Masterminds/glide) you don't, you have a `glide.yaml`: `{init,get,install}`

Simple examples
---------------

### Hello world

    package main

    import "fmt"

    func main() {
        fmt.Println("Hi world!")
    }

### Hello world 2

    package main

    func main() {
        var s string = "hi" // Not stylish enough for vet

        var s2 string
        s2 = "there"

        var s3 = "jobandtalent"

        s4 := "folks"

        println(s, s2, s3, s4)
    }

More examples
-------------

1. [Named errors](/examples/01_errors)
2. [JSON encoding/decoding](examples/02_json)
3. [Interfaces](examples/03_interfaces)
4. [Concurrency](examples/04_concurrency)
5. [Channels](examples/05_channels)
6. [APIs](examples/06_apis)

Tools
-----

- [go imports](https://godoc.org/golang.org/x/tools/cmd/goimports)
- [go fmt](https://blog.golang.org/go-fmt-your-code)
- [go vet](https://golang.org/cmd/vet/)
- [go doc](https://blog.golang.org/godoc-documenting-go-code)
