package screen

import "fmt"

type Color struct {
	R, G, B byte
}

func setFG(c Color) {
	fmt.Printf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func reset() {
	fmt.Print("\033[0m")
}
