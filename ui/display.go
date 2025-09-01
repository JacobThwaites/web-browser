package ui

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

type Window struct {
	runes   []rune
	Text    string
	Counter int
	Focused bool // whether the input box is active
}

// Input box position and size
var (
	boxX, boxY          = 50, 200
	boxWidth, boxHeight = 540, 40
)

func (g *Window) Update() error {
	// Check mouse click for focus
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= boxX && x <= boxX+boxWidth && y >= boxY && y <= boxY+boxHeight {
			g.Focused = true
		} else {
			g.Focused = false
		}
	}

	if g.Focused {
		// Add runes typed by the user
		g.runes = ebiten.AppendInputChars(g.runes[:0])
		g.Text += string(g.runes)

		// Keep at most 1 line (since it's a text box)
		ss := strings.Split(g.Text, "\n")
		if len(ss) > 1 {
			g.Text = ss[0] // only keep first line
		}

		// Enter key does nothing (ignore line breaks)
		// Backspace handling
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(g.Text) >= 1 {
				g.Text = g.Text[:len(g.Text)-1]
			}
		}
	}

	g.Counter++
	return nil
}

func (g *Window) Draw(screen *ebiten.Image) {
	// Draw the input box
	ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxWidth), float64(boxHeight), color.RGBA{200, 200, 200, 255})

	// Show text with cursor if focused
	t := g.Text
	if g.Focused && g.Counter%60 < 30 {
		t += "|"
	}
	ebitenutil.DebugPrintAt(screen, t, boxX+8, boxY+8)

	// Optional instruction text
	ebitenutil.DebugPrintAt(screen, "Click inside the box to type.", 50, 100)
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
