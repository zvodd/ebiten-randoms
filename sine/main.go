package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	clr_BgPurple    = color.RGBA{56, 52, 69, 255}
	clr_Violet      = color.RGBA{151, 85, 201, 255}
	clr_Greenish    = color.RGBA{79, 201, 109, 255}
	clr_QuiteOrange = color.RGBA{250, 180, 30, 255}

	scrDim = image.Point{1240 / 2, 680 / 2}

	texBG *ebiten.Image

	texGraph    *ebiten.Image
	graphPos    *image.Rectangle
	textStatBar *ebiten.Image
	KeySys      = NewKeyEventSys()
	toggleFlag  = false
	statbarIdx  = 0
	statbarMax  = 2
)

func init() {
	grhW := (scrDim.X) - (scrDim.X % (scrDim.Y / 2))
	grpH := (scrDim.Y / 2)
	texGraph = ebiten.NewImage(grhW, grpH)
	texGraph.Fill(clr_QuiteOrange)

	textStatBar = ebiten.NewImage(scrDim.X, 20)
	texBG = ebiten.NewImage(scrDim.X, scrDim.Y)
	texBG.Fill(clr_BgPurple)

	KeySys.AddPressHandler(ebiten.KeySpace, func(g *Game) { toggleFlag = !toggleFlag })
	KeySys.AddPressHandler(ebiten.KeyA, func(g *Game) {})
	KeySys.AddPressHandler(ebiten.KeyF1, func(g *Game) { statbarIdx = (statbarIdx + 1) % statbarMax })
}

func cycleStatusBar() {

}

// Game implements ebiten.Game interface.
type Game struct {
	frame int64
	count int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	KeySys.UpdateInput(g)
	g.frame += 1
	return nil
}

func renderSine(x int) {

	// The total render space, we use the height.
	unitpx := texGraph.Bounds().Max.Y

	//
	ymiddle := float32(float64(unitpx) / 2.0)

	// X restrained to unit length, divided by unitlength == range of  0.0 ~ 1.0
	xranged := float64(x%unitpx) / float64(unitpx)

	// X * (1 radians) == 0 ~ 2Pi || 0~360 Degrees
	sinx := math.Sin(float64(xranged) * (2 * math.Pi))

	texGraph.Set(x, int(ymiddle*float32(sinx)+ymiddle), clr_Violet)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw background
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(texBG, op)

	if toggleFlag {
		// pause
	} else {
		renderSine(g.count)
		g.count++
		if g.count > texGraph.Bounds().Max.X {
			g.count = 0
			texGraph.Fill(clr_QuiteOrange)
		}
	}

	// draw center peice
	op.GeoM.Translate(float64((scrDim.X-texGraph.Bounds().Dx())/2), float64((scrDim.Y-texGraph.Bounds().Dy())/2))
	op.CompositeMode = ebiten.CompositeModeCopy
	screen.DrawImage(texGraph, op)

	// draw bottom
	optextStatBar := &ebiten.DrawImageOptions{}
	optextStatBar.GeoM.Translate(0, float64(scrDim.Y-20))
	textStatBar.Fill(color.RGBA{50, 0, 100, 255})
	statbarMsg := ""
	switch statbarIdx {
	case 1:
		statbarMsg = fmt.Sprintf("[Frame] %-d [Graph X] %-d", g.frame, g.count)
	case 0:
		fallthrough
	default:
		statbarMsg = fmt.Sprintf("[FPS]%6.2f [TPS]%6.2f [Frame] %d", ebiten.CurrentFPS(), ebiten.CurrentTPS(), g.frame)
	}
	ebitenutil.DebugPrint(textStatBar, statbarMsg)
	screen.DrawImage(textStatBar, optextStatBar)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrDim.X, scrDim.Y
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	game := &Game{}
	// Sepcify the window size as you like. Here a 4x size is specified.
	ebiten.SetWindowSize(scrDim.X*2, scrDim.Y*2)
	ebiten.SetWindowTitle("Sine Render")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
