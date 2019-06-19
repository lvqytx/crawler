package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)              // 提交任务
	ConfigMasterWorkerChan(chan Request) // 配置初始请求任务
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v", item, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
