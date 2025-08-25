package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonState int

const (
	ButtonStateNone ButtonState = iota
	ButtonStateFocus
)

type Text struct {
	Content  string
	FontSize int32
	PosX     int32
	PosY     int32
	Color    color.RGBA
}

func (text *Text) Draw() {
	rl.DrawText(text.Content, text.PosX, text.PosY, text.FontSize, text.Color)
}

type Button struct {
	rl.Rectangle
	Text       Text
	State      ButtonState
	Color      color.RGBA
	FocusColor color.RGBA
}

func (button *Button) Draw() {
	color := button.Color
	switch button.State {
	case ButtonStateNone:
		color = button.Color
	case ButtonStateFocus:
		color = button.FocusColor
	}

	rl.DrawRectangleRec(button.Rectangle, color)
}
