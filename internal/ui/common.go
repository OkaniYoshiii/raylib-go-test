package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Content  string
	FontSize int32
	PosX     int32
	PosY     int32
	Color    color.RGBA
}

type Button struct {
	rl.Rectangle
	Color color.RGBA
}
