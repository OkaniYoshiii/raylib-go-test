package scenes

import (
	"fmt"

	"github.com/OkaniYoshiii/raylib-go-test/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	Title  ui.Text
	Text   ui.Text
	Button ui.Button
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

	return MainMenu{
		Title:  title,
		Text:   text,
		Button: button,
	}
}

func (mM *MainMenu) Update() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), mM.Button.Rectangle) {
		mM.Button.Color = rl.Black
		mM.Text.Color = rl.Red
	}
}

func (mM *MainMenu) Draw() {
	rl.ClearBackground(rl.White)
	rl.DrawText(mM.Title.Content, mM.Title.PosX, mM.Title.PosY, mM.Title.FontSize, rl.Red)

	rl.DrawRectangleRec(mM.Button.Rectangle, mM.Button.Color)
	rl.DrawText(mM.Text.Content, mM.Text.PosX, mM.Text.PosY, mM.Text.FontSize, mM.Text.Color)
}
