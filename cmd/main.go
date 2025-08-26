//Jeu d'automatisation avec un cap de productivité à passer sur chaque niveau !

package main

import (
	"github.com/OkaniYoshiii/raylib-go-test/internal/overlays"
	"github.com/OkaniYoshiii/raylib-go-test/internal/scenes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 960
const screenHeight = 540
const windowTitle = "Raylib Go Test"
const targetFPS = 60

var scene scenes.Scene
var overlay overlays.Overlay

func main() {
	// IMPORTANT: Init window must be the first Raylib function to be called
	// If not, some functions like "rl.MeasureText" won't work
	rl.InitWindow(screenWidth, screenHeight, windowTitle)
	defer rl.CloseWindow()

	rl.SetExitKey(0)

	mainMenu := scenes.NewMainMenu(screenWidth, screenHeight)
	scene = &mainMenu

	rl.SetTargetFPS(targetFPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		scene.Update()
		scene.Draw()

		// if mainMenu, ok := scene.(*scenes.MainMenu); ok {
		// if !mainMenu.Overlay.IsVisible && rl.CheckCollisionPointRec(rl.GetMousePosition(), mainMenu.Button.Rectangle()) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		// 	lvlOne := scenes.NewLevelOne(screenWidth, screenHeight)
		// 	scene = &lvlOne
		// }

		// if mainMenu.Overlay.IsVisible {
		// 	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), mainMenu.Overlay.Button.Rectangle()) {
		// 		break
		// 	}
		// }
		// }

		rl.EndDrawing()
	}
}
