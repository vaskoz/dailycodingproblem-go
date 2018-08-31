package day10

import (
	"fmt"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	t.Parallel()
	sched := NewScheduler()
	sched.Schedule(func() { fmt.Println("hi there") }, 10)
	sched.Schedule(func() { fmt.Println("hello there") }, 20)
	sched.Schedule(func() { fmt.Println("so long") }, 30)
	sched.Schedule(func() { time.Sleep(1 * time.Second); fmt.Println("i fell asleep") }, 1)
	sched.Shutdown()
}

func TestSchedulerParallel(t *testing.T) {
	t.Parallel()
	sched := NewScheduler()
	sched.Schedule(func() { fmt.Println("hi there") }, 10)
	sched.Schedule(func() { fmt.Println("hello there") }, 20)
	sched.Schedule(func() { fmt.Println("so long") }, 30)
	sched.Schedule(func() { time.Sleep(2 * time.Second); fmt.Println("i fell asleep again") }, 1)
	time.Sleep(2 * time.Second)
	sched.Schedule(func() { fmt.Println("later jobs") }, 1)
	sched.Schedule(func() { fmt.Println("later still") }, 10)
	sched.Shutdown()
}
