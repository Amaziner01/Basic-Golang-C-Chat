package main

import (
	"time"
)

type Queue struct {
	arr []*Qwarg
}
type Qwarg struct {
	mf  func(*Server, *Client)
	ma1 *Server
	ma2 *Client
}

func InitQueue() *Queue {
	return &Queue{}
}

func NewQwarg(f func(*Server, *Client), a1 *Server, a2 *Client) *Qwarg {
	return &Qwarg{mf: f, ma1: a1, ma2: a2}
}

func AddtoQueue(queue *Queue, qwarg *Qwarg) {
	//fmt.Printf("Queued...\n")
	queue.arr = append(queue.arr, qwarg)
}

//Dequeue
func Dequeue(queue *Queue) {
	if len(queue.arr) > 0 {
		qwarg := queue.arr[0]
		qwarg.mf(qwarg.ma1, qwarg.ma2)
		queue.arr = queue.arr[1:]
		//fmt.Printf("Dequeued...\n")
	}
	time.Sleep(time.Nanosecond * 5)
}
