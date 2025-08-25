package scenes

import (
	"fmt"

	"github.com/OkaniYoshiii/raylib-go-test/internal/overlays"
	"github.com/OkaniYoshiii/raylib-go-test/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	Title   ui.Text
	Text    ui.Text
	Button  ui.Button
	Overlay overlays.Menu
}

func NewMainMenu(screenWidth int32, screenHeight int32) MainMenu {
	title := func() ui.Text {
		content := "Untitled Game"
		fontSize := int32(32)
		width := rl.MeasureText(content, fontSize)

		posX, posY := screenWidth/2-width/2, screenHeight/2-fontSize/2
		fmt.Println(width, screenWidth, screenHeight, posX, posY)

		return ui.Text{
			Content:  content,
			FontSize: fontSize,
			PosX:     posX,
			PosY:     posY,
		}
	}()

	text := func() ui.Text {
		content := "Click to START"
		fontSize := int32(16)
		width := rl.MeasureText(content, fontSize)
		posX := screenWidth/2 - width/2
		posY := screenHeight/2 - fontSize/2 + title.FontSize + 5

		return ui.Text{
			Content:  content,
			FontSize: fontSize,
			PosX:     posX,
			PosY:     posY,
			Color:    rl.Black,
		}
	}()

	button := func() ui.Button {
		paddingInline := int32(12)
		paddingBlock := int32(8)
		posX := float32(text.PosX - paddingInline)
		posY := float32(text.PosY - paddingBlock)

		width := float32(rl.MeasureText(text.Content, text.FontSize) + (paddingInline * 2))
		height := float32(text.FontSize + (paddingBlock * 2))

		button := ui.Button{
			Rectangle: rl.Rectangle{
				X:      posX,
				Y:      posY,
				Width:  width,
				Height: height,
			},
			Color: rl.Red,
		}

		return button
	}()

	menuOverlay := overlays.NewMenu(screenWidth, screenHeight)

	return MainMenu{
		Title:   title,
		Text:    text,
		Button:  button,
		Overlay: menuOverlay,
	}
}

func (menu *MainMenu) Update() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), menu.Button.Rectangle) {
		menu.Button.Color = rl.Black
		menu.Text.Color = rl.Red
	}

	if menu.Overlay.IsVisible {
		isClickOutsideOfBackground := rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !rl.CheckCollisionPointRec(rl.GetMousePosition(), menu.Overlay.Background.Rectangle)
		if isClickOutsideOfBackground || rl.IsKeyPressed(rl.KeyEscape) {
			menu.Overlay.IsVisible = !menu.Overlay.IsVisible
		}
	} else if rl.IsKeyPressed(rl.KeyEscape) {
		menu.Overlay.IsVisible = true
	}
}

func (menu *MainMenu) Draw() {
	rl.ClearBackground(rl.White)
	rl.DrawText(menu.Title.Content, menu.Title.PosX, menu.Title.PosY, menu.Title.FontSize, rl.Red)

	rl.DrawRectangleRec(menu.Button.Rectangle, menu.Button.Color)
	rl.DrawText(menu.Text.Content, menu.Text.PosX, menu.Text.PosY, menu.Text.FontSize, menu.Text.Color)

	if menu.Overlay.IsVisible {
		menu.Overlay.Draw()
	}
}
