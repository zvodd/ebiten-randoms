package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"math/rand"

	vec "github.com/atedja/go-vector"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


type Ball struct{
	Pos vec.Vector
	Vel vec.Vector
}

var (
	screenX = 800
	screenY = 600
	
	// sizesList = []int{10,30,82, 128}

	bgImage               *ebiten.Image	
	infoBar          *ebiten.Image
	
	genTex *ebiten.Image

	balls []*Ball
	bOrig []Ball
)



func init() {
	infoBar = ebiten.NewImage(screenX, 20)
	bgImage = ebiten.NewImage(screenX, screenY)
	bgImage.Fill(Clr_BgPurple)
	
	tmp := GenTex(16, 2, 1, func(r, f, e, x, y int) uint8{
		b := uint8(0xff)
		if r*r < (x-r)*(x-r) + (y-r) * (y-r){
			b = 0
		}
		return b
	})
	genTex = ScaleAndColorTex(tmp, Clr_QuiteOrange, 1)


	for i := 0; i < 10; i++ {
		minX :=  float64(screenX) * 0.1
		minY :=  float64(screenY) * 0.1

		x := minX + (rand.Float64() * (float64(screenX) - minX * 2))
		y := minY + (rand.Float64() * (float64(screenY) - minY * 2))

		b := Ball{
			Pos: vec.NewWithValues([]float64{ x , y}),
			Vel: vec.NewWithValues([]float64{rand.Float64() * 10,rand.Float64() * 10}),
		}
		balls = append(balls, &b)
	}
	for _,b := range(balls){
		 bOrig = append(bOrig, *b)
	}
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	UpdateDelta()
	
	if inpututil.IsKeyJustPressed(ebiten.KeyR){
		for i, b := range(balls){
			ob := bOrig[i]
			b.Pos = ob.Pos
			// b.Vel = ob.Vel
		}
	}

	for _, b := range( balls) {
		amount := b.Vel.Clone()
		amount.Scale(1/float64(Delta))
		b.Pos = vec.Add(b.Pos, amount)
	}

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
	bottomInfoOpts := &ebiten.DrawImageOptions{}
	bottomInfoOpts.GeoM.Translate(0, float64(screenY-20))

	for _, b := range( balls){
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.Pos[0], b.Pos[1])
		screen.DrawImage(genTex,op)
	}

	infoBar.Fill(color.RGBA{50, 0, 100, 255})
	ebitenutil.DebugPrint(infoBar, fmt.Sprint("[n] next,  [space] hide/show,  Delta:", Delta,
	//  " B[0] Pos:", balls[0].Pos,
	//  " B[0] Vel:", balls[0].Vel,
	 ))
	screen.DrawImage(infoBar, bottomInfoOpts)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}


func main() {


	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenX*2, screenY*2)
	ebiten.SetWindowTitle("Textures")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}


