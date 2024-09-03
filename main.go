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
	)

	var (
		cell   rune
		vx, vy = 1, 1
		px, py int
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
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// set ball position
		board[px][py] = true
		// drawing buffer length
		// *2 for extra spaces
		// +1 for newlines
		bufLen := (width*2 + 1) * height
		// Use buffer for performance
		buf := make([]rune, 0, bufLen)
		// Slice buffer slice to 0 length.
		// This keeps the backing array the same with the len and cap
		// to ensure we're using the same buffer each time.
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
