package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonProperties struct {
	rl.Vector2

	Label    string
	FontSize int32
	Color    color.RGBA

	PaddingInline   int32
	PaddingBlock    int32
	BackgroundColor color.RGBA
}

func (props *ButtonProperties) Rectangle() rl.Rectangle {
	rect := rl.Rectangle{}

	rect.X = props.X
	rect.Y = props.Y
	rect.Width = float32((props.PaddingInline * 2) + rl.MeasureText(props.Label, props.FontSize))
	rect.Height = float32((props.PaddingBlock * 2) + props.FontSize)

	return rect
}

type Button struct {
	ButtonProperties
	FocusState[ButtonProperties]
	HoverState[ButtonProperties]
}

func (button *Button) DrawFocus() {
	if button.onFocus == nil {
		button.Draw()
		return
	}

	properties := button.onFocus(button.ButtonProperties)
	drawButton(properties)
}

func (button *Button) DrawHover() {
	if button.OnHover == nil {
		button.Draw()
		return
	}

	properties := button.OnHover(button.ButtonProperties)

	drawButton(properties)
}

func (button *Button) Draw() {
	drawButton(button.ButtonProperties)
}

func drawButton(properties ButtonProperties) {
	rl.DrawRectangleRec(properties.Rectangle(), properties.BackgroundColor)
	rl.DrawText(properties.Label, int32(properties.X)+properties.PaddingInline, int32(properties.Y)+properties.PaddingBlock, properties.FontSize, properties.Color)
}
