package main

import (
	"image"
	"image/color"

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