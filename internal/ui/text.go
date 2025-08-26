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

func (text *Text) DrawFocus() {
	properties := text.OnHover(text.TextProperties)

	text.draw(properties)
}

func (text *Text) draw(properties TextProperties) {
	rl.DrawText(properties.Content, int32(properties.X), int32(properties.Y), properties.FontSize, properties.Color)
}
