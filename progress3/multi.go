package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

const fps = 25

var (
	// Stdout defines where output gets printed to
	Stdout io.Writer = os.Stdout
	// BarFormat defines the bar design
	BarFormat = "[#>-]"
)

// MultiProgressBar is a helper for printing multiple progress bars
type MultiProgressBar struct {
	ProgressBars []*progressbar.ProgressBar

	lastPrintTime time.Time
}

// AddProgressBar adds another progress bar to the multi struct
func (mp *MultiProgressBar) AddProgressBar(p *progressbar.ProgressBar) {
	mp.ProgressBars = append(mp.ProgressBars, p)

	if len(mp.ProgressBars) > 1 {
		fmt.Println()
	}
	mp.Print()
}

// Print writes all progress bars to stdout
func (mp *MultiProgressBar) Print() {
	moveCursorUp(uint(len(mp.ProgressBars)))

	for _, p := range mp.ProgressBars {
		moveCursorDown(1)
		p.RenderBlank()
	}
}

// LazyPrint writes all progress bars to stdout if a significant update occurred
func (mp *MultiProgressBar) LazyPrint() {
	forced := false
	for _, p := range mp.ProgressBars {

		updateRequired := p.IsFinished()
		if updateRequired {
			forced = true
			break
		}
	}

	now := time.Now()
	if !forced {
		forced = now.Sub(mp.lastPrintTime) > time.Second/fps
	}

	if forced {
		mp.lastPrintTime = now

		moveCursorUp(uint(len(mp.ProgressBars)))
		for _, p := range mp.ProgressBars {
			moveCursorDown(1)
			p.RenderBlank()
		}
	}
}

/*
 * goprogressbar
 *     Copyright (c) 2016-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

func clearCurrentLine() {
	fmt.Fprintf(Stdout, "\033[2K\r")
}

func moveCursorUp(lines uint) {
	fmt.Fprintf(Stdout, "\033[%dA", lines)
}

func moveCursorDown(lines uint) {
	fmt.Fprintf(Stdout, "\033[%dB", lines)
}
