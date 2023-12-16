package main

/*
type Job3 struct {
	Name     string
	Progress int
}

func updateProgress(jobs []Job) {
	for _, job := range jobs {
		//		fmt.Printf("\r%s: %s>%s (%d%%)\n", job.Name, strings.Repeat("=", job.Progress), strings.Repeat(" ", 100-job.Progress), job.Progress)
		//fmt.Printf("\033[F%s: %s>%s (%d%%)\n", job.Name, strings.Repeat("=", job.Progress), strings.Repeat(" ", 100-job.Progress), job.Progress)
		fmt.Printf("%s: %s>%s (%d%%)\033[K\n", job.Name, strings.Repeat("=", job.Progress), strings.Repeat(" ", 100-job.Progress), job.Progress)
	}
}

func worker3(job *Job, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	defer close(ch)

	// 여기에서 실제 작업 수행
	for i := 0; i <= 100; i += 10 {
		job.Progress = i
		time.Sleep(500 * time.Millisecond) // 실제 작업 대체 필요
		ch <- struct{}{}
	}
}

func main3() {
	var wg sync.WaitGroup
	jobs := []Job{
		{Name: "Job1", Progress: 0},
		{Name: "Job2", Progress: 0},
		{Name: "Job3", Progress: 0},
		// ... 더 많은 작업 추가
	}

	for i := range jobs {
		wg.Add(1)
		ch := make(chan struct{})
		go worker(&jobs[i], &wg, ch)

		go func(job Job, ch <-chan struct{}) {
			for range ch {
				updateProgress(jobs)
			}
		}(jobs[i], ch)
	}

	wg.Wait()
	//wg.Done()

	fmt.Println() // 모든 작업이 완료된 후에 줄 바꿈
}
*/
