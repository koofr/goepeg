goepeg
======

Go epeg library bindings

## Dependencies

To install goepeg you need to install epeg C library.
Goepeg uses epeg fork that supports transformations.

https://github.com/koofr/epeg

## Install

    export CGO_CFLAGS="-I/path/to/epeg/dist/include"
    export CGO_LDFLAGS="-L/path/to/epeg/dist/lib"
    export LD_LIBRARY_PATH="-L/path/to/epeg/dist/lib"

    go get github.com/koofr/goepeg

## Testing

    go test
