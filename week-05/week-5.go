package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SlidingWindow struct {
	SucceedSlice []uint64
	FailSlice    []uint64
}

func NewSlidingWindow() *SlidingWindow {
	ret := &SlidingWindow{
		SucceedSlice: make([]uint64, 10),
		FailSlice: make([]uint64, 10),
	}
	return ret
}

func SucceedCount(window *SlidingWindow){
	index :=  time.Now().Nanosecond()/1000/1000/100
	atomic.AddUint64(&window.SucceedSlice[index], 1)
}

func FailCount(window *SlidingWindow){
	index :=  time.Now().Nanosecond()/1000/1000/100
	atomic.AddUint64(&window.FailSlice[index], 1)
}

func Avg(window *SlidingWindow) (uint64, uint64){
	return sum(window.SucceedSlice)/10, sum(window.FailSlice)/10
}

func sum(array []uint64) uint64 {
	result := uint64(0)
	for _, v := range array {
		result = result + v
	}
	return result
}

func main() {
	window := NewSlidingWindow()
	for i := 1; i <= 1000; i++ {
		time.Sleep(time.Millisecond*50)
		SucceedCount(window)
		FailCount(window)
	}
	fmt.Println(Avg(window))
}