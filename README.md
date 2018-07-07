goepeg
======

Go epeg library bindings

## Dependencies

To install goepeg you need to install epeg C library.
Goepeg uses epeg fork that supports transformations.

https://github.com/koofr/epeg

## Install

```
go get github.com/koofr/goepeg
```

## Testing

```
go test
```

## Update libepeg

```
git checkout https://github.com/koofr/epeg.git
cd epeg
git checkout feature-static
mkdir build
cd build
cmake ..
make install
```

```
cp epeg/dist/lib/libepeg.a libepeg_linux_amd64.a
cp epeg/dist/lib/libjpeg.a libjpeg_linux_amd64.a
cp epeg/dist/include/Epeg.h Epeg.h
```
