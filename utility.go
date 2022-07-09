package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func FailIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CloneRect(src *sdl.Rect) sdl.Rect {
	return sdl.Rect{
		X: src.X,
		Y: src.Y,
		W: src.W,
		H: src.H,
	}
}
