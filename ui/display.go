package ui

import (
	"image/color"
	"io"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
    screenWidth  = 640
    screenHeight = 480

    boxX, boxY          = 50, 20
    boxWidth, boxHeight = 460, 40
    btnWidth, btnHeight = 80, 40
    btnX, btnY          = boxX + boxWidth, boxY
)

type Window struct {
    runes     []rune
    Text      string
    Counter   int
    Focused   bool
    ButtonHit bool

    fontFace *text.GoTextFace
}

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

// NewWindow creates a Window, given some font source reader and size
func NewWindow(fontData io.Reader, fontSize float64) (*Window, error) {
    src, err := text.NewGoTextFaceSource(fontData)
    if err != nil {
        return nil, err
    }
    face := &text.GoTextFace{Source: src, Size: fontSize}
    return &Window{
        fontFace: face,
    }, nil
}

func (g *Window) Update() error {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()
        // input box focus
        if x >= boxX && x <= boxX+boxWidth && y >= boxY && y <= boxY+boxHeight {
            g.Focused = true
        } else {
            g.Focused = false
        }
        // button click
        if x >= btnX && x <= btnX+btnWidth && y >= btnY && y <= btnY+btnHeight {
            g.ButtonHit = true
        } else {
            g.ButtonHit = false
        }
    }

    if g.Focused {
        g.runes = ebiten.AppendInputChars(g.runes[:0])
        g.Text += string(g.runes)
        // keep only first line
        ss := strings.Split(g.Text, "\n")
        if len(ss) > 1 {
            g.Text = ss[0]
        }
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
    // Draw input box
    vector.DrawFilledRect(screen,
        float32(boxX), float32(boxY),
        float32(boxWidth), float32(boxHeight),
        color.RGBA{200, 200, 200, 255}, false,
    )

    // Draw button
    btnColor := color.RGBA{150, 150, 250, 255}
    if g.ButtonHit {
        btnColor = color.RGBA{100, 100, 200, 255}
    }
    vector.DrawFilledRect(screen,
        float32(btnX), float32(btnY),
        float32(btnWidth), float32(btnHeight),
        btnColor, false,
    )

    // Draw button label "OK"
    {
        op := &text.DrawOptions{}
        op.GeoM.Translate(float64(btnX+20), float64(btnY+28))
        op.ColorScale.ScaleWithColor(color.Black)
        text.Draw(screen, "OK", g.fontFace, op)
    }

    // Draw the user text with cursor if focused
    t := g.Text
    if g.Focused && g.Counter%60 < 30 {
        t += "|"
    }
    {
        op := &text.DrawOptions{}
        op.GeoM.Translate(float64(boxX+8), float64(boxY+28))
        op.ColorScale.ScaleWithColor(color.Black)
        text.Draw(screen, t, g.fontFace, op)
    }
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}
