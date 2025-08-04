package set

type Set struct {
	Data []int
}

func NewSet(nums ...int) *Set {
	newSet := &Set{
		Data: make([]int, 0, len(nums)),
	}

	newSet.Add(nums...)

	return newSet
}

func (s *Set) Add(nums ...int) {
	seen := make(map[int]bool)

	for _, num := range s.Data {
		seen[num] = true
	}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			s.Data = append(s.Data, num)
		}
	}
}

func (s *Set) Delete(num int) {
	for i, n := range s.Data {
		if n == num {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
		} 
	}
}

// create a Set data structure of ints with the following functionality:

// - Contains - takes an int as argument, returns true if that int is in the set, false if not

// - Intersection - takes another Set as an argument, returns a new Set that contains only the elements in the two sets, e.g. setA.Intersection(setB) => should return a new set with only elements in both setA AND setB

// - Difference - like Intersection but setA.Difference(setB) => should return only elements that are in setA but NOT in setB

// - Union - like Intersection but setA.Difference(setB) => should return all elements that are in setA AND setB (but the resulting set should not contain any duplicates)

// - String - returns a string of the elements in the set in a nice format (e.g. pretty format for logging)

// - ToList - returns a slice of ints containing the elements from the set
