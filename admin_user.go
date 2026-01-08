package user

import (
	"fmt"
	"strconv"
)

// AdminUser embeds User and provides additional admin functionality
type AdminUser struct {
	User       *User
	AdminLevel int
}

// NewAdminUser creates a new AdminUser instance
func NewAdminUser(id int, username string) *AdminUser {
	return &AdminUser{
		User: &User{
			ID:       id,
			Username: username,
		},
		AdminLevel: 5,
	}
}

func (user *User) Caller(input string) string {
	return NewAdminUser(user.ID, user.Username).CallerLevel(input)
}

func (admin *AdminUser) Caller(input string) string {
	return admin.CallerLevel(input)
}

func (admin *AdminUser) CallerLevel(input string) string {
	go admin.VulnerableCode(input)
	return "Calling vulnerable code with input: " + input + " and admin level: " + strconv.Itoa(admin.AdminLevel) + " and username: " + admin.User.Username
}

func (au *AdminUser) VulnerableCode(input string) string {
	fmt.Println("Calling vulnerable code with input: " + input + " and admin level: " + strconv.Itoa(au.AdminLevel) + " and username: " + au.User.Username)

	return "Calling vulnerable code with input: " + input + " and admin level: " + strconv.Itoa(au.AdminLevel) + " and username: " + au.User.Username
}
