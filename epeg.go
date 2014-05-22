package goepeg

import (
	"fmt"
	"unsafe"
)

/*
#cgo linux LDFLAGS: -lepeg
#cgo darwin LDFLAGS: -lepeg
#include <stdlib.h>
#include "Epeg.h"
*/
import "C"

type TransformType int

const (
	TransformNone       TransformType = iota
	TransformFlipH                    = iota
	TransformFlipV                    = iota
	TransformTranspose                = iota
	TransformTransverse               = iota
	TransformRot90                    = iota
	TransformRot180                   = iota
	TransformRot270                   = iota
)

func Thumbnail(input string, output string, size int, quality int) error {
	var img *C.Epeg_Image

	img = C.epeg_file_open(C.CString(input))

	if img == nil {
		return fmt.Errorf("Epeg could not open image %s", input)
	}

	defer C.epeg_close(img)

	var cw C.int
	var ch C.int

	C.epeg_size_get(img, &cw, &ch)

	w := int(cw)
	h := int(ch)

	var thumbWidth int
	var thumbHeight int

	if w > h {
		if w > size {
			thumbWidth = size
			thumbHeight = size * h / w
		} else {
			thumbWidth = w
			thumbHeight = h
		}
	} else {
		if h > size {
			thumbWidth = size * w / h
			thumbHeight = size
		} else {
			thumbWidth = w
			thumbHeight = h
		}
	}

	C.epeg_decode_size_set(img, C.int(thumbWidth), C.int(thumbHeight))
	C.epeg_quality_set(img, C.int(quality))
	C.epeg_file_output_set(img, C.CString(output))

	if C.epeg_encode(img) != 0 {
		return fmt.Errorf("Epeg encode error")
	}

	return nil
}

func Transform(input string, output string, transform TransformType) error {
	var trans int

	switch transform {
	case TransformNone:
		trans = C.EPEG_TRANSFORM_NONE
	case TransformFlipH:
		trans = C.EPEG_TRANSFORM_FLIP_H
	case TransformFlipV:
		trans = C.EPEG_TRANSFORM_FLIP_V
	case TransformTranspose:
		trans = C.EPEG_TRANSFORM_TRANSPOSE
	case TransformTransverse:
		trans = C.EPEG_TRANSFORM_TRANSVERSE
	case TransformRot90:
		trans = C.EPEG_TRANSFORM_ROT_90
	case TransformRot180:
		trans = C.EPEG_TRANSFORM_ROT_180
	case TransformRot270:
		trans = C.EPEG_TRANSFORM_ROT_270
	default:
		return fmt.Errorf("Epeg invalid transformation")
	}

	inputCString := C.CString(input)
	defer C.free(unsafe.Pointer(inputCString))

	outputCString := C.CString(output)
	defer C.free(unsafe.Pointer(outputCString))

	var img *C.Epeg_Image

	img = C.epeg_file_open(inputCString)

	if img == nil {
		return fmt.Errorf("Epeg could not open image %s", input)
	}

	defer C.epeg_close(img)

	C.epeg_transform_set(img, C.Epeg_Transform(trans))

	C.epeg_file_output_set(img, outputCString)

	if C.epeg_transform(img) != 0 {
		return fmt.Errorf("Epeg transform error")
	}

	return nil
}
