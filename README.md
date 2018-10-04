goepeg
======

Go wrapper for epeg.

[Epeg](https://github.com/mattes/epeg) is a C library for insanely fast JPEG/JPG thumbnail scaling with the minimum fuss and CPU overhead. It makes use of libjpeg features of being able to load an image by only decoding the DCT coefficients needed to reconstruct an image of the size desired.

The library has bundled archive which are automatically linked. This means that installing libepeg and libjpeg is not necessary.

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
git clone https://github.com/koofr/epeg.git
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
