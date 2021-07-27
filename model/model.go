package model

import "image"

type ImgParam struct {
	FileName  string
	NewWidth  uint
	NewHeight uint
	Rectangle image.Rectangle
}
