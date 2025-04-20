package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const appVersion = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	arg := os.Args[1]

	switch arg {
	case "-h", "--help":
		showHelp()
		return
	case "-v", "version":
		showVersion()
		return
	default:
		password := arg
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error hashing password:", err)
			return
		}
		fmt.Println("Hashed password:", string(hashedPassword))
	}
}

func showHelp() {
	fmt.Println("gbh: Gerundium bcrypt hasher")
	fmt.Println("Usage:")
	fmt.Println("  go run gbh <password>")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  go run gbh mySecurePassword123")
	fmt.Println("  go run gbh -h          # Show help")
	fmt.Println("  go run gbh -v          # Show version")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help    Show this help message")
	fmt.Println("  -v, --version    Show application version")
}

func showVersion() {
	fmt.Println("Gerundium bcrypt hasher version:", appVersion)
}