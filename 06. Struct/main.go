package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Id          int
	Rating      float64
}

func main() {
	user := NewUser(
		"Johnny",
		28,
		"79212905591",
		4.5,
		false,
		123,
	)
	fmt.Println(user)
	user.RatingUp(4)
	fmt.Println(user.GetRating())
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	rating float64,
	isClose bool,
	id int,
) User {
	if age < 0 || age > 150 {
		return User{}
	}
	if name == "" {
		return User{}
	}
	if phoneNumber == "" {
		return User{}
	}
	if id <= 0 {
		return User{}
	}
	if rating < 0.0 || rating > 10.0 {
		return User{}
	}
	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
		Id:          id,
	}
}

func (user *User) RatingUp(rating float64) {
	newRating := user.Rating + rating
	if newRating < 10.0 && newRating > 0.0 {
		user.Rating += rating
	} else {
		fmt.Println("Не валидный рейтинг")
	}

}

func (user *User) GetRating() float64 {
	return user.Rating
}
