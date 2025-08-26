package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Component interface {
	Rectangle() rl.Rectangle
	Draw()
}

type Focusable interface {
	Component
	DrawFocus()
}

type Hoverable interface {
	Component
	DrawHover()
}

type UserInterface struct {
	Elements       []Component
	FocusedElement Focusable
	HoveredElement Hoverable
}

func (ui *UserInterface) Update() {
	ui.HoveredElement = nil

	for i := range len(ui.Elements) {
		isHovered := rl.CheckCollisionPointRec(rl.GetMousePosition(), ui.Elements[i].Rectangle())

		if element, ok := ui.Elements[i].(Hoverable); ok && isHovered {
			ui.HoveredElement = element
		}

		if element, ok := ui.Elements[i].(Focusable); ok {
			if ui.FocusedElement == element || isHovered && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				ui.FocusedElement = element
			}
		}
	}
}

func (ui *UserInterface) Draw() {
	for i := range len(ui.Elements) {
		if element, ok := ui.Elements[i].(Focusable); ok && element == ui.FocusedElement {
			element.DrawFocus()
			continue
		}

		if element, ok := ui.Elements[i].(Hoverable); ok && element == ui.HoveredElement {
			element.DrawHover()
			continue
		}

		ui.Elements[i].Draw()
	}
}
