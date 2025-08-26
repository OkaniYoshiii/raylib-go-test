package scenes

import (
	"image/color"

	"github.com/OkaniYoshiii/raylib-go-test/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	// Title  ui.Text
	UI ui.UserInterface
	// Overlay overlays.Menu
}

func NewMainMenu(screenWidth int32, screenHeight int32) MainMenu {
	title := ui.Text{}
	title.Content = "UntitledGame"
	title.Color = rl.Red
	title.FontSize = 24
	title.X = float32((screenWidth / 2) - (rl.MeasureText(title.Content, title.FontSize) / 2))
	title.Y = float32((screenHeight / 2) - (title.FontSize / 2))

	button := ui.Button{
		ButtonProperties: ui.ButtonProperties{
			Vector2: rl.Vector2{
				X: 0,
				Y: 0,
			},
			Label:           "Click to START",
			PaddingInline:   12,
			PaddingBlock:    8,
			FontSize:        16,
			Color:           rl.Black,
			BackgroundColor: rl.Red,
		},
	}

	button.OnHover = func(state ui.ButtonProperties) ui.ButtonProperties {
		state.Color = color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}

		return state
	}

	button.OnFocus(func(state ui.ButtonProperties) ui.ButtonProperties {
		state.BackgroundColor.G = 150

		return state
	})

	buttonTwo := ui.Button{
		ButtonProperties: ui.ButtonProperties{
			Vector2: rl.Vector2{
				X: 150,
				Y: 0,
			},
			Label:           "Button Two",
			PaddingInline:   12,
			PaddingBlock:    8,
			FontSize:        16,
			Color:           rl.Black,
			BackgroundColor: rl.Red,
		},
	}

	return MainMenu{
		UI: ui.UserInterface{
			Elements: []ui.Component{
				&button,
				&buttonTwo,
				&title,
			},
		},
	}
}

func (menu *MainMenu) Update() {
	menu.UI.Update()
}

func (menu *MainMenu) Draw() {
	rl.ClearBackground(rl.White)

	menu.UI.Draw()
}
