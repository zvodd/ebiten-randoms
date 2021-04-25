package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


var (
	screenX = 800
	screenY = 600
	
	// sizesList = []int{10,30,82, 128}

	bgImage               *ebiten.Image	
	infoBar          *ebiten.Image
	
	genTex *ebiten.Image
	tfuncs = []TexGenFunc{TexFuncQuad, TexFunc1, TexFunc2, TexFunc3, TexFunc4}
	tindx = 0
	
	preview = true
	tic = 1
	maxtic=1280
)



func init() {
	infoBar = ebiten.NewImage(screenX, 20)
	bgImage = ebiten.NewImage(screenX, screenY)
	bgImage.Fill(Clr_BgPurple)
	
	tmp := GenTex(64, 4, 1, TexFuncQuad)
	genTex = ScaleAndColorTex(tmp, Clr_QuiteOrange, 1)
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	tic++
	if tic > maxtic{ tic =1}

	if ebiten.IsKeyPressed(ebiten.KeyR){
		bgImage.Fill(Clr_BgPurple)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace){
		preview = !preview
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN){
		tindx++
		if tindx > len(tfuncs)-1{ tindx = 0}

	}

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
	bottomInfoOpts := &ebiten.DrawImageOptions{}
	bottomInfoOpts.GeoM.Translate(0, float64(screenY-20))

	
	if preview{
		tmp := GenTex(64, 2, tic, tfuncs[tindx])
		genTex = ScaleAndColorTex(tmp, Clr_QuiteOrange, 1)
		for i := 0; i < 3; i++{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(screenX)/2 -float64(genTex.Bounds().Max.X/2) + float64((i-1) * genTex.Bounds().Max.X),
				float64(screenY)/2 -float64(genTex.Bounds().Max.Y/2),
			)
			screen.DrawImage(genTex,op)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(screenX)/2 -float64(genTex.Bounds().Max.X/2) + float64((i-1) * genTex.Bounds().Max.X),
				float64(screenY)/2 +float64(genTex.Bounds().Max.Y/2),
			)
			screen.DrawImage(genTex,op)
		}
	}

	infoBar.Fill(color.RGBA{50, 0, 100, 255})
	ebitenutil.DebugPrint(infoBar, fmt.Sprint("[n] next,  [space] hide/show,  index:", tindx, ", f:", tic) )
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


