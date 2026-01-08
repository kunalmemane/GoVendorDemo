package main

import (
	"fmt"
	"log"

	"github.com/kunalmemane/GoVendorDemo/user"
)

func main() {
	fmt.Println("=== User Package Demo ===\n")

	// Create a regular user
	regularUser := user.NewUser(1, "john_doe", "password123", "john@example.com", "user")
	fmt.Printf("Created User: %s (ID: %d, Role: %s)\n", regularUser.Username, regularUser.ID, regularUser.Role)

	// Demonstrate vulnerable code
	fmt.Println("\n--- Vulnerable Code Example ---")
	vulnerableInput := "password' OR '1'='1"
	vulnerableQuery := regularUser.VulnerableCode(vulnerableInput)
	fmt.Printf("Vulnerable Query: %s\n", vulnerableQuery)
	fmt.Println("⚠️  WARNING: This is vulnerable to SQL injection!")

	// Demonstrate non-vulnerable code
	fmt.Println("\n--- Non-Vulnerable Code Example ---")
	safeInput := "password123"
	safeQuery, err := regularUser.NonVulnerableCode(safeInput)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Safe Query: %s\n", safeQuery)
		fmt.Println("✅ This is safe - input is validated and sanitized")
	}

	// Test with malicious input in non-vulnerable method
	fmt.Println("\n--- Non-Vulnerable Code with Malicious Input ---")
	maliciousInput := "password' OR '1'='1"
	safeQuery2, err := regularUser.NonVulnerableCode(maliciousInput)
	if err != nil {
		fmt.Printf("Error caught: %v\n", err)
	} else {
		fmt.Printf("Safe Query (sanitized): %s\n", safeQuery2)
		fmt.Println("✅ Malicious characters were removed")
	}

	// Create an admin user
	fmt.Println("\n\n=== Admin User Demo ===\n")
	adminUser := user.NewAdminUser(2, "admin_user", "admin123", "admin@example.com", 5)
	fmt.Printf("Created AdminUser: %s (ID: %d, AdminLevel: %d)\n", 
		adminUser.User.Username, adminUser.User.ID, adminUser.AdminLevel)
	fmt.Printf("Permissions: %v\n", adminUser.Permissions)

	// Demonstrate vulnerable code for admin
	fmt.Println("\n--- Admin Vulnerable Code Example ---")
	adminVulnerableInput := "admin' OR '1'='1' --"
	adminVulnerableQuery := adminUser.VulnerableCode(adminVulnerableInput)
	fmt.Printf("Vulnerable Admin Query: %s\n", adminVulnerableQuery)
	fmt.Println("⚠️  WARNING: This is extremely dangerous for admin users!")

	// Demonstrate non-vulnerable code for admin
	fmt.Println("\n--- Admin Non-Vulnerable Code Example ---")
	adminSafeInput := "admin123"
	adminSafeQuery, err := adminUser.NonVulnerableCode(adminSafeInput)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Safe Admin Query: %s\n", adminSafeQuery)
		fmt.Println("✅ This is safe - input is validated, sanitized, and admin level is checked")
	}

	// Test with malicious input in admin non-vulnerable method
	fmt.Println("\n--- Admin Non-Vulnerable Code with Malicious Input ---")
	adminMaliciousInput := "admin' OR '1'='1' /*"
	adminSafeQuery2, err := adminUser.NonVulnerableCode(adminMaliciousInput)
	if err != nil {
		fmt.Printf("Error caught: %v\n", err)
	} else {
		fmt.Printf("Safe Admin Query (sanitized): %s\n", adminSafeQuery2)
		fmt.Println("✅ Malicious characters were removed")
	}

	// Test with empty input
	fmt.Println("\n--- Testing Empty Input Validation ---")
	emptyResult, err := regularUser.NonVulnerableCode("")
	if err != nil {
		fmt.Printf("✅ Validation worked: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", emptyResult)
	}

	fmt.Println("\n=== Demo Complete ===")
}
