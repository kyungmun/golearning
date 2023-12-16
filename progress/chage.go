package main

import (
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
)

type Job struct {
	name     string
	progress *pb.ProgressBar
}

func worker(job *Job, wg *sync.WaitGroup) {
	defer wg.Done()

	//progress := job.progress
	job.progress.Set("prefix", "["+job.name+"]")
	job.progress.SetTemplateString(`{{string . "prefix"}}{{bar . "[" "=" ">" "-" "]"}} {{percent . }} {{string . "suffix"}}`)
	//job.progress.SetTemplateString("template", `{{string . "prefix"}}{{counters . }} {{bar . "[" "=" ">" "-" "]"}} {{percent . }} {{string . "suffix"}}`)

	//job.progress.SetWidth(80)
	// 여기에서 실제 작업 수행
	for i := 0; i < 100; i += 1 {
		if job.name == "Job1" {
			time.Sleep(10 * time.Millisecond) // 실제 작업 대체 필요
		}
		if job.name == "Job2" {
			time.Sleep(30 * time.Millisecond) // 실제 작업 대체 필요
		}

		time.Sleep(100 * time.Millisecond) // 실제 작업 대체 필요
		job.progress.Increment()
	}
	job.progress.Finish()
}

func main() {
	var wg sync.WaitGroup
	jobs := []Job{
		{name: "Job1"},
		{name: "Job2"},
		{name: "Job3"},
	}

	// 진행률 표시 준비
	first := pb.New(100)
	first.SetMaxWidth(100)
	second := pb.New(100)
	second.SetMaxWidth(100)
	third := pb.New(100)
	third.SetMaxWidth(100)
	jobs[0].progress = first
	jobs[1].progress = second
	jobs[2].progress = third

	// pool, err := pb.StartPool(first, second, third)
	// if err != nil {
	// 	panic(err)
	// }

	pool := pb.NewPool()
	pool.Add(first)
	pool.Add(second)
	pool.Add(third)
	pool.Start()

	for i := range jobs {
		wg.Add(1)
		go worker(&jobs[i], &wg)
	}

	wg.Wait()
	pool.Stop()

	// 작업 완료 후 터미널을 깨끗하게 하기 위해 끝에 개행 추가
	println()
}
