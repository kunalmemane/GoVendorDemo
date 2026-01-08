package main

import (
	"fmt"
	"strings"
	"time"
)

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

// GetFullName returns the full name of the person
func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// GetEmail returns the email address
func (p *Person) GetEmail() string {
	return p.Email
}

// SetEmail updates the email address
func (p *Person) SetEmail(email string) {
	p.Email = email
}

// IsAdult checks if the person is an adult (age >= 18)
func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

// Activate sets the person's active status to true
func (p *Person) Activate() {
	p.IsActive = true
}

// Deactivate sets the person's active status to false
func (p *Person) Deactivate() {
	p.IsActive = false
}

// ValidateEmail checks if the email contains @ symbol
func (p *Person) ValidateEmail() bool {
	return strings.Contains(p.Email, "@")
}

// GetAge returns the person's age
func (p *Person) GetAge() int {
	return p.Age
}

// UpdateAge updates the person's age
func (p *Person) UpdateAge(newAge int) {
	if newAge >= 0 {
		p.Age = newAge
	}
}

// String returns a string representation of the Person
func (p *Person) String() string {
	status := "inactive"
	if p.IsActive {
		status = "active"
	}
	return fmt.Sprintf("Person{ID: %d, Name: %s %s, Email: %s, Age: %d, Status: %s}",
		p.ID, p.FirstName, p.LastName, p.Email, p.Age, status)
}
