package overlays

import (
	"image/color"

	"github.com/OkaniYoshiii/raylib-go-test/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	rl.Rectangle

	Color color.RGBA
}

type Menu struct {
	Button     ui.Button
	Text       ui.Text
	Background Background
	IsVisible  bool
}

func NewMenu(screenWidth int32, screenHeight int32) Menu {
	text := func() ui.Text {
		text := ui.Text{}

		text.Content = "Quit game"
		text.FontSize = 16
		text.Color = rl.Red

		width := rl.MeasureText(text.Content, text.FontSize)

		text.PosX = screenWidth/2 - width/2
		text.PosY = screenHeight/2 - width/2

		return text
	}()

	button := func() ui.Button {
		button := ui.Button{}

		padding := int32(12)

		button.Width = float32(rl.MeasureText(text.Content, text.FontSize) + (padding * 2))
		button.Height = float32(text.FontSize + (padding * 2))

		button.X = float32(text.PosX - padding)
		button.Y = float32(text.PosY - padding)

		button.Color = rl.Black

		return button
	}()

	background := func() Background {
		background := Background{}

		screenSpace := float32(85.0 / 100.0)

		background.X = float32(screenWidth) * ((1 - screenSpace) / 2)
		background.Y = float32(screenHeight) * ((1 - screenSpace) / 2)

		background.Width = float32(screenWidth) * screenSpace
		background.Height = float32(screenHeight) * screenSpace

		background.Color = color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 200,
		}

		return background
	}()

	return Menu{
		Button:     button,
		Text:       text,
		Background: background,
	}
}

func (menu *Menu) Draw() {
	rl.DrawRectangleRec(menu.Background.Rectangle, menu.Background.Color)
	rl.DrawRectangleRec(menu.Button.Rectangle, menu.Button.Color)
	rl.DrawText(menu.Text.Content, menu.Text.PosX, menu.Text.PosY, menu.Text.FontSize, menu.Text.Color)
}
