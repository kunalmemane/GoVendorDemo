package user

import (
	"fmt"
	"strings"
)

// VulnerableCode demonstrates a vulnerable method with security issues
// VULNERABLE: Directly concatenates user input without sanitization
// This can lead to SQL injection, XSS, or other injection attacks
func (u *User) VulnerableCode(input string) string {
	// VULNERABLE: No input validation or sanitization
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s'", u.Username, input)
	return query
}

// NonVulnerableCode demonstrates a safe method with proper validation
// SAFE: Validates and sanitizes input before use
func (u *User) NonVulnerableCode(input string) (string, error) {
	// SAFE: Input validation
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	// SAFE: Sanitize input - remove potentially dangerous characters
	sanitized := strings.ReplaceAll(input, "'", "")
	sanitized = strings.ReplaceAll(sanitized, "\"", "")
	sanitized = strings.ReplaceAll(sanitized, ";", "")
	sanitized = strings.ReplaceAll(sanitized, "--", "")

	// SAFE: Use parameterized approach (simulated)
	query := fmt.Sprintf("SELECT * FROM users WHERE username = ? AND password = ?")
	// In real implementation, use parameterized queries with placeholders

	return fmt.Sprintf("%s (username: %s, sanitized_input: %s)", query, u.Username, sanitized), nil
}
