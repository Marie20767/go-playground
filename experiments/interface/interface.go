package interfaces

import (
	"fmt"
)

type meercat struct {
	name string
	age int
}

type elephant struct {
	name string
}

type animal interface {
	greet() string
	kind() string
}

func main() {

	meercat := meercat{
		name: "Alfred",
		age: 2,
	}

	fmt.Println(doGreet(meercat))


	elephant := elephant{
		name: "Kongi",
	}

	fmt.Println(doGreet(elephant))
}

func doGreet(a animal) string {
		return a.greet()
}

// func getKind (a animal) string {
// 	return a.kind()
// }

func (m meercat) greet() string {
	return fmt.Sprintf("Holita, I am a %s. My name is %s.", m.kind(), m.name)
}

func (m meercat) kind() string {
	return "meercat"
}

func (e elephant) greet() string {
	return fmt.Sprintf("Helloooo, I am a %s. My name is %s.", e.kind(), e.name)
}

func (e elephant) kind() string {
	return "elephant"
}