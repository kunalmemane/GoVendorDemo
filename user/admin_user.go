package user

import (
	"fmt"
	"strings"
)

// AdminUser embeds User and provides additional admin functionality
type AdminUser struct {
	User        *User
	AdminLevel  int
	Permissions []string
}

// NewAdminUser creates a new AdminUser instance
func NewAdminUser(id int, username, password, email string, adminLevel int) *AdminUser {
	return &AdminUser{
		User: &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
			Role:     "admin",
			IsActive: true,
		},
		AdminLevel:  adminLevel,
		Permissions: []string{"read", "write", "delete"},
	}
}

// VulnerableCode demonstrates a vulnerable method with security issues
// VULNERABLE: Directly concatenates user input without sanitization
// This can lead to SQL injection, XSS, or other injection attacks
func (au *AdminUser) VulnerableCode(input string) string {
	// VULNERABLE: No input validation or sanitization
	// Even more dangerous for admin users with elevated privileges
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s' AND admin_level = %d",
		au.User.Username, input, au.AdminLevel)
	return query
}

// NonVulnerableCode demonstrates a safe method with proper validation
// SAFE: Validates and sanitizes input before use
func (au *AdminUser) NonVulnerableCode(input string) (string, error) {
	// SAFE: Input validation
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	// SAFE: Additional validation for admin operations
	if au.AdminLevel < 1 {
		return "", fmt.Errorf("insufficient admin level")
	}

	// SAFE: Sanitize input - remove potentially dangerous characters
	sanitized := strings.ReplaceAll(input, "'", "")
	sanitized = strings.ReplaceAll(sanitized, "\"", "")
	sanitized = strings.ReplaceAll(sanitized, ";", "")
	sanitized = strings.ReplaceAll(sanitized, "--", "")
	sanitized = strings.ReplaceAll(sanitized, "/*", "")
	sanitized = strings.ReplaceAll(sanitized, "*/", "")

	// SAFE: Use parameterized approach (simulated)
	query := fmt.Sprintf("SELECT * FROM users WHERE username = ? AND password = ? AND admin_level = ?")
	// In real implementation, use parameterized queries with placeholders

	return fmt.Sprintf("%s (username: %s, admin_level: %d, sanitized_input: %s)",
		query, au.User.Username, au.AdminLevel, sanitized), nil
}
