package main

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	/*
	   		mpb := MultiProgressBar{}

	   		var wg sync.WaitGroup

	   		for i := 0; i < 2; i++ {
	   			pb := progressbar.Default(100)
	   			mpb.AddProgressBar(pb)
	   		}

	   		for i := 0; i < len(mpb.ProgressBars); i++ {
	   			wg.Add(1)
	   			go worker(mpb.ProgressBars[i], nil, &mpb, &wg)
	   		}

	   		wg.Wait()
	   		//fmt.Println()
	   	}

	   func worker(pb *progressbar.ProgressBar, tpb *progressbar.ProgressBar, mpb *MultiProgressBar, wg *sync.WaitGroup) {

	   	//pb.Total = 150
	   	// fill progress bars one after another
	   	for i := 1; i <= int(pb.GetMax()); i++ {
	   		pb.Add(1)
	   		//tpb.Current++

	   		//mpb.LazyPrint()

	   		time.Sleep(23 * time.Millisecond)
	   		// if pb.Text == "Progress 1" {
	   		// 	time.Sleep(23 * time.Millisecond)
	   		// }
	   	}
	   	wg.Done()
	*/
}
