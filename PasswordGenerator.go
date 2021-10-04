package main

import (
	"fmt"
	"math/rand"
	"time"
)

// All letters, numbers and symbols that can be used to generate a password
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"!@#$%^&*()-_=+"

// A new random seed to make different random passwords
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().Unix()))

// Function that creates a random generated String
func String(length int, charset string) string {
	// New empty slice of bytes with the length given with length parameter
	b := make([]byte, length)
	// Loop that adds a random character to the slive until it meets the given length
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// variables for the length and count of the passwords
	var length, count int

	fmt.Println("How many characters should the password have?")
	fmt.Scanln(&length)
	fmt.Println("How many passwords should be generated?")
	fmt.Scanln(&count)

	// counts the number given off the user for creating a number of passwords
	for i := 0; i <= count; i++ {
		fmt.Println(String(length, charset))
	}
}
