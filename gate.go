package main

type Gate struct {
	QueueNums   int
	QueueLength int
	// 8个队列，用channel
	Queues []chan *Message
}

func NewGate(n int) *Gate {
	g := new(Gate)
	g.QueueNums = n
	g.QueueLength = 100
	g.Queues = make([]chan *Message, 0)
	for i := 0; i < n; i++ {
		g.Queues = append(g.Queues, make(chan *Message, g.QueueLength))
	}
	return g
}
