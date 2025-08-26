package ui

import (
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Counter struct {
	rl.Vector2

	Label    string
	FontSize int32
	Color    color.RGBA
	Count    int
}

func (counter *Counter) Draw() {
	label := counter.Label + " : " + strconv.Itoa(counter.Count)
	rl.DrawText(label, int32(counter.X), int32(counter.Y), counter.FontSize, counter.Color)
}
