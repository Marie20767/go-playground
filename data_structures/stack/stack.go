package stack

type Stack struct {
	data []rune
}

func NewStack() *Stack {
	return &Stack{
		data: make([]rune, 0),
	}
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	topIndex := len(s.data) - 1

	top := s.data[topIndex]
	s.data = s.data[:topIndex]

	return top, true
}

func (s *Stack) Push(el rune) {
	s.data = append(s.data, el)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
