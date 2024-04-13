package stack

type Stack[T any] struct {
	Elem []T
}

func (s *Stack[T]) Push(item T) error {
	s.Elem = append(s.Elem, item)
	return nil
}

func (s *Stack[T]) Pop() (T, bool) {
	var res T
	if len(s.Elem) <= 0 {
		return res, false
	}
	res = s.Elem[len(s.Elem)-1]
	s.Elem = s.Elem[:len(s.Elem)-1]
	return res, true
}
