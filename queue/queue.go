package main

import "errors"

const MAX_ELEM_SIZE = 10
type Queue struct{
	head int
	tail int
	len int
	data [MAX_ELEM_SIZE]int
}

func(q *Queue) New() *Queue {
	q.len = 0
	q.tail = -1
	q.head = 0
	return q
}

func(q* Queue) Enqueu(data int) error{
	if q.len >= MAX_ELEM_SIZE{
		return errors.New("Queue is full")
	}
	q.tail = (q.tail + 1) % MAX_ELEM_SIZE
	q.data[q.tail] = data
	q.len++
	return nil
}

func(q* Queue) Dequeu(data int) (int,error){
	if(q.len == 0){
		return 0,errors.New("Queue is empty")
	}
	item := q.data[q.head]
	q.head = (q.head + 1) % MAX_ELEM_SIZE
	q.len--

	return item,nil
}

func main()  {
	var q = Queue{}
	q.New()
	q.Enqueu(10)
	q.Enqueu(100)
}