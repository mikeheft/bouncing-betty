package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {
	const (
		cellEmpty = ' '
		cellFull  = '⚾'
		maxFrames = 1200
		speed     = time.Second / 20
		// initial velocities
		ivx, ivy = 5, 2
	)

	var (
		cell     rune
		vx, vy   = ivx, ivx // velocities
		px, py   int        // ball position
		ppx, ppy int        // previous ball position
	)

	// Get size of screen dynamically
	width, height := screen.Size()
	// get the rune width of the ball emoji
	ballWidth := runewidth.RuneWidth(cellFull)
	// adjust the width and height
	width /= ballWidth

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	// Set max time for animation
	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		// Redirect ball
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		if py <= 0 || py >= height-ivx {
			vy *= -1
		}

		// remove the previous ball and put the new ball
		board[px][py], board[ppx][ppy] = true, false

		// save the previous positions
		ppx, ppy = px, py

		// set ball position
		board[px][py] = true

		// drawing buffer length
		// *2 for extra spaces
		// +1 for newlines
		bufLen := (width*2 + 1) * height

		// Use buffer for performance
		buf := make([]rune, 0, bufLen)

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellFull
				}

				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))
		time.Sleep(speed)
	}
}
