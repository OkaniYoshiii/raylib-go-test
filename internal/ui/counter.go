package ui

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Counter struct {
	Label Text
	Count int

	count Text
}

func NewCounter(x int32, y int32, label string, fontSize int32) Counter {
	counter := Counter{}

	counter.Label = func() Text {
		text := Text{}

		text.Content = label
		text.FontSize = fontSize
		text.PosX, text.PosY = x, y

		text.Color.A = 255

		return text
	}()

	counter.count = func() Text {
		text := Text{}

		text.FontSize = fontSize
		text.PosX = counter.Label.PosX + rl.MeasureText(counter.Label.Content, counter.Label.FontSize)
		text.PosY = counter.Label.PosY

		text.Color.A = 255

		return text
	}()

	return counter
}

func (counter *Counter) Draw() {
	counter.Label.Draw()

	counter.count.Content = " : " + strconv.Itoa(counter.Count)
	counter.count.Draw()
}
