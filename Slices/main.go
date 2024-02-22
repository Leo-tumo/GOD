package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
	"time"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		// initial velocities
		ivx, ivy = 1, 1
	)

	var (
		px, py int
		vx, vy = ivx, ivy

		cell rune
	)

	width, height := screen.Size()
	width /= runewidth.RuneWidth(cellBall)
	height--

	buf := make([]rune, 0, (width*2+1)*height)

	screen.Clear()
	fmt.Print("\033[?25l")

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		buf = buf[:0]

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cell = cellEmpty

				if px == x && py == y {
					cell = cellBall
				}

				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print("[1;31m", string(buf))

		time.Sleep(speed)
	}
	fmt.Print("\033[?25h")
}
