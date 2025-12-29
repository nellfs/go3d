package main

import "github.com/nellfs/go3d/internal/screen"

func main() {
	width := 13
	height := 7
	canvas := screen.NewCanvas(width, height, '█', screen.Color{84, 84, 92})

	canvas.Clear()
	canvas.PutPixel(width/2, height/2, '█', screen.Color{48, 120, 255})
	canvas.Render()
}
