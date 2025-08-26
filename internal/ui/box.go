package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoxProperties struct {
	rl.Vector2
	PaddingInline   int32
	PaddingBlock    int32
	BackgroundColor color.RGBA
}

func (properties *BoxProperties) Rectangle() rl.Rectangle {
	return rl.Rectangle{
		X:      properties.X,
		Y:      properties.Y,
		Width:  float32(properties.PaddingInline),
		Height: float32(properties.PaddingBlock),
	}
}

type Box struct {
	BoxProperties
	FocusState[BoxProperties]
	HoverState[BoxProperties]
}

func (box *Box) draw() {
	rl.DrawRectangleRec(box.Rectangle(), box.BackgroundColor)
}
