package user

// User represents a user with credentials and profile information
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Role     string
	IsActive bool
}

// NewUser creates a new User instance
func NewUser(id int, username, password, email, role string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
		Role:     role,
		IsActive: true,
	}
}
