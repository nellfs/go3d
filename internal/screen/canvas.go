package screen

import (
	"fmt"
)

type pixel struct {
	Char  rune
	Color Color
}

type canvas struct {
	Width      int
	Height     int
	Cells      [][]pixel
	Background rune
	Color      Color
}

func NewCanvas(width, height int, background rune, color Color) *canvas {
	c := &canvas{
		Width:      width,
		Height:     height,
		Cells:      make([][]pixel, height),
		Background: background,
		Color:      color,
	}

	for y := 0; y < height; y++ {
		c.Cells[y] = make([]pixel, width)
		for x := 0; x < width; x++ {
			c.Cells[y][x] = pixel{Char: background, Color: color}
		}
	}

	return c
}

func (c *canvas) Clear() {
	for y := 0; y < len(c.Cells); y++ {
		for x := 0; x < len(c.Cells[y]); x++ {
			c.Cells[y][x] = pixel{Char: c.Background, Color: c.Color}
		}
	}
}

func (c *canvas) PutPixel(x, y int, character rune, color Color) {
	sx := c.Width/2 + x
	sy := c.Height/2 + y
	c.Cells[sy][sx] = pixel{Char: character, Color: color}
}

func (c *canvas) Render() {
	for y := 0; y < len(c.Cells); y++ {
		for x := 0; x < len(c.Cells[y]); x++ {
			cell := c.Cells[y][x]
			setFG(cell.Color)
			fmt.Print(string(cell.Char))
			reset()
		}
		fmt.Println()
	}
}
