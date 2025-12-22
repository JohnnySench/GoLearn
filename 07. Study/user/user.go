package user

type User struct {
	age  int
	name string
}

func NewUser(age int, name string) User {
	if validateName(name) {
		return User{}
	}
	if validateAge(age) {
		return User{}
	}
	return User{
		age:  age,
		name: name,
	}
}

func (user *User) SetAge(age int) {
	if age <= 150 && age > 0 {
		user.age = age
	}
}

func (user *User) SetName(name string) {
	if name != "" {
		user.name = name
	}
}

func validateName(name string) bool {
	return name == ""
}

func validateAge(age int) bool {
	return age <= 0 || age >= 150
}
