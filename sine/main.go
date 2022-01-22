package main

import (
	"fmt"
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

	screenX = 320
	screenY = 280

	bgImage *ebiten.Image
	imgopts = &ebiten.DrawImageOptions{}

	texGraph   *ebiten.Image
	infoBar    *ebiten.Image
	KeySys     = NewKeyEventSys()
	toggleFlag = false
)

func init() {
	texGraph = ebiten.NewImage(screenX/2, screenY/2)
	texGraph.Fill(clr_QuiteOrange)
	infoBar = ebiten.NewImage(screenX, 20)
	bgImage = ebiten.NewImage(screenX, screenY)
	bgImage.Fill(clr_BgPurple)

	KeySys.AddPressHandler(ebiten.KeySpace, func(g *Game) { toggleFlag = !toggleFlag })
	KeySys.AddPressHandler(ebiten.KeyA, func(g *Game) {})
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

	unitpx := texGraph.Bounds().Max.Y

	ymiddle := float32(float64(unitpx) / 2.0)
	xranged := float64(x%unitpx) / float64(unitpx)     // X restrained to unit length, divided by unitlength == range of  0.0 ~ 1.0
	sinx := math.Sin(float64(xranged) * (2 * math.Pi)) // X * (1 radians) == 0 ~ 2Pi || 0~360 Degrees

	texGraph.Set(x, int(ymiddle*float32(sinx)+ymiddle), clr_Violet)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw background
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(bgImage, op)

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
	op.GeoM.Translate(float64(screenX/4), float64(screenY/4))
	op.CompositeMode = ebiten.CompositeModeCopy
	screen.DrawImage(texGraph, op)

	// draw bottom
	bottomInfoOpts := &ebiten.DrawImageOptions{}
	bottomInfoOpts.GeoM.Translate(0, float64(screenY-20))
	infoBar.Fill(color.RGBA{50, 0, 100, 255})
	ebitenutil.DebugPrint(infoBar, fmt.Sprint("Frame:", g.frame, ", Graph idx:", g.count))
	screen.DrawImage(infoBar, bottomInfoOpts)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
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
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenX*4, screenY*4)
	ebiten.SetWindowTitle("Paint Toy")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
