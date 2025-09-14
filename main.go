package main

import (
	"bytes"
	"log"
	"os"

	"web-browser/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
    fontBytes, err := os.ReadFile("assets/DejaVuSans.ttf")
    if err != nil {
        log.Fatalf("failed to read font: %v", err)
    }

    g, err := ui.NewWindow(bytes.NewReader(fontBytes), 16)
    if err != nil {
        log.Fatalf("failed to create window: %v", err)
    }

    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Web Browser")

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
