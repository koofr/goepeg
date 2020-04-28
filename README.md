goepeg
======

Go wrapper for epeg.

[Epeg](https://github.com/mattes/epeg) is a C library for insanely fast JPEG/JPG thumbnail scaling with the minimum fuss and CPU overhead. It makes use of libjpeg features of being able to load an image by only decoding the DCT coefficients needed to reconstruct an image of the size desired.

The library has bundled archive which are automatically linked. This means that installing libepeg and libjpeg is not necessary.

Current archive files are built from `epeg` `v0.1.0` (`libjpeg-turbo` `2.0.4`).

## Install

```
go get github.com/koofr/goepeg
```

## Testing

```
go test
```

## Update libepeg

```sh
git clone https://github.com/koofr/epeg.git
```

### macOS

```sh
cd epeg
mkdir build
cd build
cmake ..
make install
```

```sh
cp epeg/dist/lib/libepeg.a libepeg_darwin_amd64.a
cp epeg/dist/lib/libjpeg.a libjpeg_darwin_amd64.a
```

### Linux

Docker is needed to build for Ubuntu 12.04 compatibility:

```sh
cd epeg
docker build -t epeg .
docker run -d --name epeg-copy epeg sh -c 'sleep 9999999'
mkdir out
cd out
docker cp epeg-copy:/epeg/dist/lib/libepeg.a libepeg_linux_amd64.a
docker cp epeg-copy:/epeg/dist/lib/libjpeg.a libjpeg_linux_amd64.a
docker cp epeg-copy:/epeg/dist/include/Epeg.h Epeg.h
docker kill epeg-copy && docker rm epeg-copy
```

```sh
cp epeg/out/{libepeg_linux_amd64.a,libjpeg_linux_amd64.a,Epeg.h} .
```
