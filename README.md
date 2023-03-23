goepeg
======

Go wrapper for epeg.

[Epeg](https://github.com/mattes/epeg) is a C library for insanely fast JPEG/JPG
thumbnail scaling with the minimum fuss and CPU overhead. It makes use of
libjpeg features of being able to load an image by only decoding the DCT
coefficients needed to reconstruct an image of the size desired.

The library has bundled archive which are automatically linked. This means that
installing libepeg and libjpeg is not necessary.

Current archive files are built from `epeg` `v0.1.0` and `libjpeg-turbo`
`2.1.91`.

## Install

```sh
go get github.com/koofr/goepeg
```

## Testing

```sh
go test
```

### Test in Docker

```sh
docker run --rm -v $(pwd):/app -w /app golang:1.20.2-bullseye sh -c 'go test'
```

## Update libepeg

```sh
git clone https://github.com/koofr/epeg.git
```

### macOS

```sh
brew install cmake nasm

cd epeg
mkdir build
cd build
cmake -DCMAKE_OSX_DEPLOYMENT_TARGET=10.14.6 ..
make install
```

```sh
cp epeg/dist/lib/libepeg.a libepeg_darwin_amd64.a
cp epeg/dist/lib/libjpeg.a libjpeg_darwin_amd64.a
cp epeg/dist/include/Epeg.h Epeg.h
```

### Linux

Docker is needed to build for Ubuntu 16.04 compatibility:

```sh
cd epeg
docker build -t epeg .
docker create --name epeg-copy epeg
mkdir out
cd out
docker cp epeg-copy:/epeg/dist/lib/libepeg.a libepeg_linux_amd64.a
docker cp epeg-copy:/epeg/dist/lib/libjpeg.a libjpeg_linux_amd64.a
docker cp epeg-copy:/epeg/dist/include/Epeg.h Epeg.h
docker rm epeg-copy
```

```sh
cp epeg/out/{libepeg_linux_amd64.a,libjpeg_linux_amd64.a,Epeg.h} .
```
