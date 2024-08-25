package queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	data []T
	size int
}

func New[T any]() *Queue[T] {
	q := new(Queue[T])
	q.data = make([]T, 0)
	q.size = 0
	return q
}

func (q *Queue[T]) Push(d T) {
	q.data = append(q.data, d)
	q.size++
}

func (q *Queue[T]) Front() (T, error) {
	if q == nil || q.size == 0 {
		var empty T
		return empty, errors.New("Empty queue")
	}
	return q.data[0], nil

}

func (q *Queue[T]) Pop() (T, error) {
	if q == nil || q.size == 0 {
		var empty T
		return empty, errors.New("Empty queue")
	}
	p := q.data[0]
	q.data = q.data[1:]
	return p, nil
}

func (q *Queue[T]) Print() {
	if q == nil || len(q.data) == 0 {
		fmt.Println("queue is empty")
		return
	}

	for _, data := range q.data {
		fmt.Println(data)
	}
}
