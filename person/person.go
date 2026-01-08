package person

import "time"

// Person represents a person with various attributes
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Age       int
	CreatedAt time.Time
	IsActive  bool
}

// NewPerson creates a new Person instance
func NewPerson(id int, firstName, lastName, email string, age int) *Person {
	return &Person{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Age:       age,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
}
