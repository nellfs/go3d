package main

import "github.com/nellfs/go3d/internal/screen"

func main() {
	width := 23
	height := width / 2
	canvas := screen.NewCanvas(width, height, '.', screen.Color{84, 84, 92})

	canvas.Clear()
	canvas.PutPixel(0, 0, '█', screen.Color{48, 120, 255})
	canvas.Render()
}
