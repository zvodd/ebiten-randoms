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

	clr_BgPurple = color.RGBA{56, 52, 69, 255}
	clr_Violet = color.RGBA{151, 85, 201, 255}
	clr_Greenish = color.RGBA{79, 201, 109, 255}
	clr_QuiteOrange = color.RGBA{250, 180, 30, 255}
	

	screenX = 1200
	screenY = 800


	bgImage               *ebiten.Image
	imgopts                = &ebiten.DrawImageOptions{}
	
	imgBrushes *ebiten.Image
	showBrushes = false

	infoBar          *ebiten.Image
	penSizes = []int{10,30,82, 128}
	penImg *ebiten.Image
	pens []*ebiten.Image
	penKeys = []ebiten.Key{ebiten.Key1,ebiten.Key2,ebiten.Key3,ebiten.Key4,ebiten.Key5,ebiten.Key6,ebiten.Key7,ebiten.Key8,ebiten.Key9}

	someSprite        image.Rectangle
	spriteList        []image.Rectangle
	
	anyKeyInteraction bool = false
	clickList         []int
	first = true

	
	spotLightImage *ebiten.Image
)



func init() {
	imgBrushes = ebiten.NewImage(screenX, 200)
	infoBar = ebiten.NewImage(screenX, 20)
	bgImage = ebiten.NewImage(screenX, screenY)
	bgImage.Fill(clr_BgPurple)
	
	
	pens =  [](*ebiten.Image){
	}
	const r = 256
	const rsqr = r*r
	// alphas := image.Point{r * 2, r * 2}
	alphas := image.Point{r * 8, r * 4}

	a := image.NewAlpha(image.Rectangle{image.Point{}, alphas})
	for j := 0; j < alphas.Y; j++ {
		for i := 0; i < alphas.X; i++ {
			// d is the distance between (i, j) and the (circle) center.
			d := math.Sqrt(float64((i-r)*(i-r) + (j-r)*(j-r)))
			// Alphas around the center are 0 and values outside of the circle are 0xff.
			
			
			
			// if d <= r { b = 0xff}
			// if d > r { b = 0}
			
			// d := float64((i-r)*(i-r) + (j-r)*(j-r))
			// b := uint8(0xff)
			//  b := uint8(0b111)
			//b := uint8(max(0, min(255, int(3*255*d/r)-510)))
			//b := uint8(max(0, min(255, int(3*d*255/r)-2*255)))

			// if d > rsqr { b = (uint8(d) - 0xff)}
			
			// b := (uint8(d-32) ^ uint8(63)) //+ uint8(0xdd)
			// if d > r { b = 0}
			
			
			// b := (uint8(d*d) - 0xdd) ^ 0xff
			b := uint8(d*d)
			// if d > r { b = 0}
			
			a.SetAlpha(i, j, color.Alpha{b})
		}
	}
	spotLightImage = ebiten.NewImageFromImage(a)

	for _, size := range(penSizes){
		// p.Fill(color.Black)
		img := ebiten.NewImage(size, size)
		// img.Fill(color.RGBA{0,0,0,255})
		img.Fill(clr_QuiteOrange)
		op := &ebiten.DrawImageOptions{}
		op.CompositeMode = ebiten.CompositeModeDestinationIn
		scl :=  float64(size)/float64(r*2)
		op.GeoM.Scale(float64(scl), float64(scl))
		img.DrawImage(spotLightImage, op)
		pens = append(pens, img)
	}
	penImg = pens[0]


}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, _ := ebiten.CursorPosition()
		clickList = append(clickList, x)
	}
	
	showBrushes = ebiten.IsKeyPressed(ebiten.KeyP)

	for i, k := range(penKeys){
		if ebiten.IsKeyPressed(k){
			if i < len(pens){
				penImg = pens[i]
			}
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyR){
		bgImage.Fill(clr_BgPurple)
	}


	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	
	
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		penops := &ebiten.DrawImageOptions{}
		penops.GeoM.Translate(float64(x), float64(y))
		// penops.CompositeMode = ebiten.CompositeModeCopy
		bgImage.DrawImage(penImg, penops)
	}

	screen.DrawImage(bgImage, imgopts)

	bottomInfoOpts := &ebiten.DrawImageOptions{}
	bottomInfoOpts.GeoM.Translate(0, float64(screenY-20))


	if showBrushes{
		imgBrushes.Fill(color.Black)
		for i, p := range(pens){
			op := &ebiten.DrawImageOptions{}
			prv := 0
			for j, ps := range(penSizes){ if j <  i {prv += ps}}
			op.GeoM.Translate(float64(/*penSizes[i]*/ prv + (i * 10)  + 20),0)
			imgBrushes.DrawImage(p, op)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0,0)
		screen.DrawImage(imgBrushes, op)
	}
	
	if ebiten.IsKeyPressed(ebiten.KeySpace){
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(spotLightImage,op)
	}

	infoBar.Fill(color.RGBA{50, 0, 100, 255})
	ebitenutil.DebugPrint(infoBar, fmt.Sprint("Info text:", clickList) )
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
	ebiten.SetWindowSize(screenX, (screenY))
	ebiten.SetWindowTitle("Paint Toy")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}


