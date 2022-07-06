# Basic Go

-   go test -bench=.
-   `go test -cover`
-   `go test ./... -cover`
-   In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
-   go install github.com/kisielk/errcheck@latest
    -   package to help identify errors that have not been checked
    -   `errcheck .`

# Resources

-   [How to Write Go Code (with GOPATH)](https://go.dev/doc/gopath_code)
-   [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
-   [Pointers in Go](https://www.golang-book.com/books/intro/8)
