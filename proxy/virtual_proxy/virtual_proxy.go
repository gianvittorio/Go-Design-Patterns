package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from: ", filename)

	return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image: ", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (lb *LazyBitmap) Draw() {
	if (lb.bitmap == nil) {
		lb.bitmap = NewBitmap(lb.filename)
	}

	lb.bitmap.Draw()
}

func main() {
	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
}
