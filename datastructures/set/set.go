package set

import (
	"strconv"
	"strings"
)

type Set struct {
	Data map[int]struct{}
}

func NewSet(nums ...int) *Set {
	newSet := &Set{
		Data: make(map[int]struct{}),
	}

	newSet.Add(nums...)

	return newSet
}

func (s *Set) Add(nums ...int) {
	for _, num := range nums {
		s.Data[num] = struct{}{}
	}
}

func (s *Set) Delete(num int) {
	delete(s.Data, num)
}

func (s *Set) Contains(num int) bool {
	_, exists := s.Data[num]; return exists
}

func (s *Set) Intersection(other *Set) *Set {
	intersection := NewSet()

	for num := range other.Data {
		if _, exists := s.Data[num]; exists {
			intersection.Data[num] = struct{}{}
		}
	}

	return intersection
}

func (s *Set) Difference(other *Set) *Set {
	difference := NewSet()

	for num := range s.Data {
		if _, exists := other.Data[num]; !exists {
			difference.Data[num] = struct{}{}
		}
	}

	return difference
}

func (s *Set) Union(other *Set) *Set {
	union := NewSet()

	for num := range s.Data {
		union.Add(num)
	}

	for num := range other.Data {
		union.Add(num)
	}

	return union
}

func (s *Set) ToList() []int {
	slice := make([]int, 0, len(s.Data))

	for num := range s.Data {
		slice = append(slice, num)
	}

	return slice
}

func (s *Set) String() string {
	keys := make([]string, 0, len(s.Data))

	for num := range s.Data {
		keys = append(keys, strconv.Itoa(num))
	}

	return strings.Join(keys, ", ")
}
