package day10

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Scheduler represents a job scheduler.
type Scheduler struct {
	queue     chan job
	wg        sync.WaitGroup
	startTime time.Time
}

type job struct {
	f            func()
	n            uint
	jobStartTime time.Time
}

// NewScheduler returns a new instance of a running job scheduler.
func NewScheduler() *Scheduler {
	s := &Scheduler{queue: make(chan job), startTime: time.Now()}
	s.wg.Add(1)
	go func() {
		for j := range s.queue {
			s.wg.Add(1)
			go func(j job) {
				ms, _ := time.ParseDuration(fmt.Sprintf("%dms", j.n)) // nolint: gosec
				time.Sleep(ms)
				log.Printf("starting job at %dms", convertNanoToMillis(time.Since(s.startTime).Nanoseconds()))
				j.f()
				log.Printf("completed job at %dms", convertNanoToMillis(time.Since(s.startTime).Nanoseconds()))
				s.wg.Done()
			}(j)
		}
	}()
	return s
}

// Schedule schedules a new job to be run.
func (s *Scheduler) Schedule(f func(), millis uint) {
	current := convertNanoToMillis(time.Since(s.startTime).Nanoseconds())
	log.Printf("Received at %dms scheduled for %dms", current, current+int64(millis))
	s.queue <- job{f, millis, time.Now()}
}

// Shutdown prevents new jobs from being scheduled and allows the scheduler to terminate.
func (s *Scheduler) Shutdown() {
	close(s.queue)
	s.wg.Done()
	s.wg.Wait()
	log.Printf("Scheduler shutdown complete at %dms", convertNanoToMillis(time.Since(s.startTime).Nanoseconds()))
}

func convertNanoToMillis(duration int64) int64 {
	return duration / int64(1000000)
}
