package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type TexGenFunc func(r, fact, e, x, y int) uint8 

func GenTex(r, sizefact, extra int, method TexGenFunc) *ebiten.Image{
	alphas := image.Point{r * sizefact, r * sizefact}
	a := image.NewAlpha(image.Rectangle{image.Point{}, alphas})
	for j := 0; j < alphas.Y; j++ {
		for i := 0; i < alphas.X; i++ {
			a.SetAlpha(i, j, color.Alpha{method(r, sizefact, extra, i, j)})
		}
	}
	return ebiten.NewImageFromImage(a)
}


func ScaleAndColorTex(src *ebiten.Image, color color.RGBA, scale int) *ebiten.Image{
	size := src.Bounds().Max
	img := ebiten.NewImage(size.X * scale, size.Y * scale)
	
	img.Fill(color)

	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeDestinationIn
	op.GeoM.Scale(float64(scale), float64(scale))
	
	img.DrawImage(src, op)
	return img
}

func TexFuncQuad(r,f,e,x,y int) uint8{
	d := math.Sqrt(float64((x-r)*(x-r) + (y-r)*(y-r))) 
	t := e - 32
	b := (uint8(d*d) - uint8(t)) ^ 0xff
	return b
}

func TexFunc1(r,f,e,x,y int) uint8{
	d := math.Sqrt(float64((x-r)*(x-r) + (y-r)*(y-r))) 
	t := e - 32
	b := (uint8(d) - uint8(t)) ^ 0xff

	return b
}
func TexFunc2(r,f,e,x,y int) uint8{

	d := math.Sqrt(float64((x-r)*(x-r) + (y-r)*(y-r))) 
	b := uint8(d*d / float64(e/r))
	return b
}
func TexFunc3(r,f,e,x,y int) uint8{

	d := (x-r)*(x-r) + (y-r)*(y-r)

	b := (uint8(d-32) ^ uint8(63)) //+ uint8(0xdd)
	return b
}
func TexFunc4(r,f,e,x,y int) uint8{

	d := math.Sqrt(float64((x-r)*(x-r) + (y-r)*(y-r)))
	b := (uint8(d-32) ^ uint8(63)) + uint8(0xdd)
	return b
}
	// d is the distance between (x, y) and the (cxrcle) center.
	// d := math.Sqrt(float64((x-r)*(x-r) + (y-r)*(y-r)))
	// Alphas around the center are 0 and values outside of the circle are 0xff.

	// if d <= r { b = 0xff}
	// if d > r { b = 0}

	// d := float64((x-r)*(x-r) + (y-r)*(y-r))
	// b := uint8(0xff)
	//  b := uint8(0b111)
	//b := uint8(max(0, min(255, int(3*255*d/r)-510)))
	//b := uint8(max(0, min(255, int(3*d*255/r)-2*255)))

	// if d > rsqr { b = (uint8(d) - 0xff)}

	// b := (uint8(d-32) ^ uint8(63)) //+ uint8(0xdd)
	// if d > r { b = 0}

	// b := (uint8(d*d) - 0xdd) ^ 0xff
	// b := uint8(d * d)
	// if d > r { b = 0}
