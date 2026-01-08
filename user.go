package user

import "fmt"

type User struct {
	ID       int
	Username string
}

func (u *User) VulnerableCode(input string) string {
	return "Calling vulnerable code with input: " + input + " and username: " + u.Username
}

func (u *User) NonVulnerableCode(input string) string {
	return "Calling non-vulnerable code with input: " + input + " and username: " + u.Username
}

func (u *User) SetID(id int) {
	u.ID = id
	fmt.Println("ID set to: ", u.ID)
}
