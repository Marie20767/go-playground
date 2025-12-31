package graphs

import "fmt"

func GraphExamples() {
	data := [][]int{
		{1, 1, 0},
		{0, 1, 1},
		{0, 1, 2},
	}

	rottingOranges := OrangesRotting(data)
	fmt.Println(">>> rottingOranges: ", rottingOranges)
}
