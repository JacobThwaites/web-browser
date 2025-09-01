package main

import (
	"log"
	"web-browser/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	screenWidth  := 640
	screenHeight := 480
	g := &ui.Window{
		Text:    "",
		Counter: 0,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Web Browser")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}