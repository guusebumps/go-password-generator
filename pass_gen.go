package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	letters   = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	specials  = "!@#$%^&*()-_=+<>?"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generatePassword(length int, useNumbers, useSpecials, useUppercase bool) (string, error) {
	charset := letters

	if useNumbers {
		charset += numbers
	}
	if useSpecials {
		charset += specials
	}
	if useUppercase {
		charset += uppercase
	}

	password := make([]byte, length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	length := flag.Int("length", 12, "Length of the password")
	useNumbers := flag.Bool("numbers", true, "Include numbers")
	useSpecials := flag.Bool("specials", true, "Include special characters")
	useUppercase := flag.Bool("uppercase", true, "Include uppercase letters")
	flag.Parse()

	password, err := generatePassword(*length, *useNumbers, *useSpecials, *useUppercase)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}
