package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {
	users := []User{
		{
			name:  "Alice",
			email: "alice_95@gmail.com",
			age:   35,
		},
		{
			name:  "Marti",
			email: "marti_kerstens@gmail.com",
			age:   29,
		},
	}

	userMapResult := getUserMap(users)

	fmt.Println("userMapResult", userMapResult)

	userSliceResult := getUserSlice(userMapResult)

	fmt.Println("userSliceResult", userSliceResult)
}

func getUserMap(users []User) map[string]User {
	userMap := make(map[string]User, len(users))

	for _, user := range users {
		userMap[user.email] = user
	}

	return userMap
}

func getUserSlice(users map[string]User) []User {
	slice := make([]User, 0, len(users))

	for _, value := range users {
		slice = append(slice, value)
	}

	return slice
}
