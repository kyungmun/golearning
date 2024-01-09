package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/muesli/goprogressbar"
)

func main2() {
	mpb := goprogressbar.MultiProgressBar{}

	goprogressbar.BarFormat = "[=>-]"
	for i := 0; i < 10; i++ {
		pb := &goprogressbar.ProgressBar{
			Text:    "Progress " + strconv.FormatInt(int64(i+1), 10),
			Total:   100,
			Current: 0,
			Width:   60,
		}

		mpb.AddProgressBar(pb)
	}

	pb := &goprogressbar.ProgressBar{
		Text:    "Overall Progress",
		Total:   1000,
		Current: 0,
		Width:   100,
	}
	mpb.AddProgressBar(pb)

	// fill progress bars one after another
	for j := 0; j < 10; j++ {
		for i := 1; i <= 100; i++ {
			p := mpb.ProgressBars[j]
			p.Current = int64(i)
			p.Text = fmt.Sprintf("%d of %d", i, p.Total)

			pb.Current++

			mpb.LazyPrint()
			time.Sleep(23 * time.Millisecond)
		}
	}

	fmt.Println()
}
