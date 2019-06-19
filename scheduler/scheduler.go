package scheduler

import (
	"crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}
