package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextProperties struct {
	rl.Vector2

	Content  string
	FontSize int32
	Color    color.RGBA
}

type Text struct {
	TextProperties
	HoverState[TextProperties]
}

func (text *Text) Rectangle() rl.Rectangle {
	rect := rl.Rectangle{}

	rect.X = text.X
	rect.Y = text.Y
	rect.Width = float32(rl.MeasureText(text.Content, text.FontSize))
	rect.Height = float32(text.FontSize)

	return rect
}

func (text *Text) DrawFocus() {
	properties := text.OnHover(text.TextProperties)

	text.draw(properties)
}

func (text *Text) Draw() {
	text.draw(text.TextProperties)
}

func (text *Text) draw(properties TextProperties) {
	rl.DrawText(properties.Content, int32(properties.X), int32(properties.Y), properties.FontSize, properties.Color)
}
