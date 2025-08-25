package ui

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ButtonsCount int = 2
)

type InventoryBarButton int

const (
	EmptyButton InventoryBarButton = iota
	HomeButton
)

type InventoryBar struct {
	rl.Rectangle

	Color       color.RGBA
	EmptyButton Button
	HomeButton  Button

	Buttons [ButtonsCount]Button
}

func NewInventoryBar(screenWidth int, screenHeight int) InventoryBar {
	inventoryBar := InventoryBar{}

	padding := 10
	gap := 10
	buttonSize := 50

	inventoryBar.Width = float32((cap(inventoryBar.Buttons) * buttonSize) + ((cap(inventoryBar.Buttons) - 1) * gap) + (padding * 2))
	inventoryBar.Height = float32(buttonSize + (padding * 2))

	inventoryBar.X = float32(screenWidth)/2 - inventoryBar.Width/2
	inventoryBar.Y = float32(screenHeight) - inventoryBar.Height - 15

	inventoryBar.Color = color.RGBA{
		A: 200,
	}

	inventoryBar.Buttons = func() [ButtonsCount]Button {
		buttons := [ButtonsCount]Button{}

		y := inventoryBar.Y + float32(padding)

		for i := range cap(buttons) {
			button := Button{}

			button.X = inventoryBar.X + float32(padding) + ((float32(buttonSize + gap)) * float32(i))
			button.Y = y

			button.Width = float32(buttonSize)
			button.Height = float32(buttonSize)

			buttons[i] = button
		}

		buttons[EmptyButton].Color = color.RGBA{
			A: 150,
		}

		buttons[EmptyButton].FocusColor = func() color.RGBA {
			color := buttons[EmptyButton].Color
			color.A = 255
			return color
		}()

		buttons[HomeButton].Color = color.RGBA{
			G: 255,
			A: 150,
		}

		buttons[HomeButton].FocusColor = func() color.RGBA {
			color := buttons[HomeButton].Color
			color.A = 255
			return color
		}()

		return buttons
	}()

	return inventoryBar
}

func (inventoryBar *InventoryBar) Draw() {
	rl.DrawRectangleRec(inventoryBar.Rectangle, inventoryBar.Color)

	for _, button := range inventoryBar.Buttons {
		button.Draw()
	}
}
