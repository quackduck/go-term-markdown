package fooimage

import (
	"image"
	"io"
)

func DecodeConfig(r io.Reader) (image.Config, string, error) {
	return image.DecodeConfig(r)
}
