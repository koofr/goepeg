goepeg
======

Go epeg library bindings

## Dependencies

To install goepeg you need to install epeg C library.
Goepeg uses epeg fork that supports transformations.

https://github.com/koofr/epeg

## Install

    go get github.com/koofr/goepeg

## Testing

To run tests you have to append `/usr/local/lib` to your `LD_LIBRARY_PATH`.

    export LD_LIBRARY_PATH=/usr/local/lib
    go test goepeg
