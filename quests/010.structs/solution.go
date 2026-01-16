package structs

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func NewUser(id int, name, email string, age int) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func (u *User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) DisplayName() string {
	return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func (u *User) Birthday() {
	u.Age++
}

func CloneUser(u User) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}
