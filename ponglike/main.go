package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"math/rand"

	vec "github.com/atedja/go-vector"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Ball struct {
	Pos vec.Vector
	Vel vec.Vector
}

var (
	screenX = 800
	screenY = 600

	// sizesList = []int{10,30,82, 128}

	bgImage *ebiten.Image
	infoBar *ebiten.Image

	genTex *ebiten.Image

	balls     []*Ball
	bOrig     []Ball
	initSpeed = float64(4)

	physFactor = 1.0 / 100.0
)

func init() {
	infoBar = ebiten.NewImage(screenX, 20)
	bgImage = ebiten.NewImage(screenX, screenY)
	bgImage.Fill(Clr_BgPurple)

	// generate a circle
	tmp := GenTex(16, 2, 1, func(r, f, e, x, y int) uint8 {
		b := uint8(0xff)
		if r*r < (x-r)*(x-r)+(y-r)*(y-r) {
			b = 0
		}
		return b
	})
	genTex = ScaleAndColorTex(tmp, Clr_QuiteOrange, 1)
	// placerBasic()
	placerPartitionShuffle()
}

// shuffle grid indexs
func placerPartitionShuffle() {
	// set partition size some division of  scrren space
	spdiv := 10
	border := 16
	W, H := (screenX-border*2)/spdiv, (screenY-border*2)/spdiv
	size := W * H
	for i, place := range Shuffled(size) {
		if i > 9 {
			break
		}
		// covert index in "flatspace" back to grid space
		X := float64(((place % W) * spdiv))
		Y := float64(((place / H) * spdiv))
		X, Y = X+float64(border), Y+float64(border)
		vx := (rand.Float64() - 0.5) * 2 * initSpeed
		vy := (rand.Float64() - 0.5) * 2 * initSpeed
		b := Ball{
			Pos: vec.NewWithValues([]float64{X, Y}),
			Vel: vec.NewWithValues([]float64{vx, vy}),
		}
		balls = append(balls, &b)
	}
	for _, b := range balls {
		bOrig = append(bOrig, *b)
	}
}

func placerBasic() {
	maxProxSqr := math.Pow(50, 2)

	for i := 0; i < 10; i++ {
		minX := float64(screenX) * 0.1
		minY := float64(screenY) * 0.1

		x := minX + (rand.Float64() * (float64(screenX) - minX*2))
		y := minY + (rand.Float64() * (float64(screenY) - minY*2))

		// check distance from previous basic
		drop := false
		for _, other := range balls {
			difx := x - other.Pos[0]
			dify := x - other.Pos[1]
			sqdis := math.Pow(difx, 2) + math.Pow(dify, 2)
			if sqdis <= maxProxSqr {
				// balls left off the table
				fmt.Printf("%d to close %.2f,%.2f, to %v  [ Distance %.2f]\n", i, x, y, other.Pos, math.Sqrt(sqdis))
				drop = true
				continue
			}
		}
		if drop {
			continue
		}

		vx := (rand.Float64() - 0.5) * 2 * initSpeed
		vy := (rand.Float64() - 0.5) * 2 * initSpeed
		b := Ball{
			Pos: vec.NewWithValues([]float64{x, y}),
			Vel: vec.NewWithValues([]float64{vx, vy}),
		}
		balls = append(balls, &b)
	}

	for _, b := range balls {
		bOrig = append(bOrig, *b)
	}
}

// Game implements ebiten.Game interface.
type Game struct {
	deltatimer *DeltaTimer
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	delta := g.deltatimer.Update()
	deltaRecip := 1.0 / delta

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for i, b := range balls {
			ob := bOrig[i]
			b.Pos = ob.Pos
			b.Vel = ob.Vel
		}
	}

	// add attraction
	for _, b := range balls {
		for _, b2 := range balls {
			if b2 == b {
				continue
			}
			dif := vec.Subtract(b2.Pos, b.Pos)
			dif.Scale(deltaRecip * physFactor)
			b.Vel = vec.Add(b.Vel, dif)

		}
	}

	// resolve, and warp across screen edges
	for _, b := range balls {
		amount := b.Vel.Clone()
		amount.Scale(physFactor)
		nextPos := vec.Add(b.Pos, amount)

		if nextPos[0] > float64(screenX) {
			nextPos[0] = 0
		} else if nextPos[0] < 0 {
			nextPos[0] = float64(screenX)
		}
		if nextPos[1] > float64(screenY) {
			nextPos[1] = 0
		} else if nextPos[1] < 0 {
			nextPos[1] = float64(screenY)
		}

		b.Pos = nextPos
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
	bottomInfoOpts := &ebiten.DrawImageOptions{}
	bottomInfoOpts.GeoM.Translate(0, float64(screenY-20))

	for _, b := range balls {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.Pos[0], b.Pos[1])
		screen.DrawImage(genTex, op)
	}

	infoBar.Fill(color.RGBA{50, 0, 100, 255})
	ebitenutil.DebugPrint(infoBar, fmt.Sprint("[r] reset,  Delta:", g.deltatimer.Delta)) //  " B[0] Pos:", balls[0].Pos,
	//  " B[0] Vel:", balls[0].Vel,

	screen.DrawImage(infoBar, bottomInfoOpts)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {

	game := &Game{deltatimer: &DeltaTimer{}}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenX*2, screenY*2)
	ebiten.SetWindowTitle("Space Chaos Pendulum")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
