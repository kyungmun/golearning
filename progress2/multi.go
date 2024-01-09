/*
 * goprogressbar
 *     Copyright (c) 2016-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/muesli/goprogressbar"
)

func main() {
	goprogressbar.BarFormat = "[=>-]"

	mpb := goprogressbar.MultiProgressBar{}

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		pb := &goprogressbar.ProgressBar{
			Text:        "Progress " + strconv.FormatInt(int64(i+1), 10),
			PrependText: "",
			Total:       100,
			Current:     0,
			Width:       100,
			// PrependTextFunc: func(p *goprogressbar.ProgressBar) string {
			// 	return fmt.Sprintf("%d of %d", p.Current, p.Total)
			// },
		}
		mpb.AddProgressBar(pb)
	}

	sumpb := &goprogressbar.ProgressBar{
		Text:        "Overall Progress",
		PrependText: "",
		Total:       300,
		Current:     0,
		Width:       100,
	}
	mpb.AddProgressBar(sumpb)

	for i := 0; i < len(mpb.ProgressBars)-1; i++ {
		wg.Add(1)
		go worker(mpb.ProgressBars[i], sumpb, &mpb, &wg)
	}

	wg.Wait()
	fmt.Println()
}

func worker(pb *goprogressbar.ProgressBar, tpb *goprogressbar.ProgressBar, mpb *goprogressbar.MultiProgressBar, wg *sync.WaitGroup) {

	pb.Total = 150
	// fill progress bars one after another
	for i := 1; i <= int(pb.Total); i++ {
		pb.Current = int64(i)
		tpb.Current++

		mpb.LazyPrint()

		time.Sleep(23 * time.Millisecond)
		if pb.Text == "Progress 1" {
			time.Sleep(23 * time.Millisecond)
		}
	}
	wg.Done()

}
