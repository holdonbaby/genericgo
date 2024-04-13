package queue

import (
	"fmt"
)

type Queue[T any] struct {
	Elem []T
}

func (q *Queue[T]) Push(item T) error {
	q.Elem = append(q.Elem, item)
	return nil
}

func (q *Queue[T]) pop() (T, bool) {
	var res T
	if len(q.Elem) <= 0 {
		return res, false
	}
	res = q.Elem[0]
	q.Elem = q.Elem[1:len(q.Elem)]

	return res, true
}

// HttpRequest
func (q *Queue[T]) HttpRequest() {
	// process the request
	request, bl := q.pop()
	if !bl {
		return
	}
	process(request)
	q.HttpRequest()
}

func process[T any](t T) {
	fmt.Printf("百度百科~~~~~~~~~~~~~~~~ %v \n", t)
}
