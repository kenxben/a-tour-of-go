package main

import(
	//"golang.org/x/tour/pic"
	"image"
	"image/color"
	"image/png"
	"os"
	//"fmt"
)

type Image struct{
	x int
	y int
	w int
	h int
	col color.Model
}

func (im Image) Bounds() image.Rectangle {
	//fmt.Println(im.x, im.y, im.w, im.h)
	return image.Rect(im.x, im.y, im.w, im.h)
}

func (im Image) ColorModel() color.Model {
	return im.col
}

func (im Image) At(x, y int) color.Color {
	//v := uint8((x*x + y*y + 7*x*y/4)/2)
	v := uint8((x*x + x*x*y + x*y*y + y*y)/200)
	return color.RGBA{v/25, 200*v, v, 255}
}

func main() {
	m := Image{x:200, y:200,w:450,h:450}
	//pic.ShowImage(m)
	
	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, m)
}

