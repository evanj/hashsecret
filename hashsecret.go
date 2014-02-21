package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func generatePhoneNumbers() chan string {
	outputChannel := make(chan string)
	go func() {
		const maxDigits = 10
		maxNumber := int64(math.Pow10(maxDigits))
		const firstNumber int64 = 2000000000

		for i := firstNumber; i < maxNumber; i++ {
			outputChannel <- strconv.FormatInt(i, 10)
		}

		close(outputChannel)
	}()

	return outputChannel
}

func main() {
	sharedSalt := "012345678"
	sharedSaltBytes := []byte(sharedSalt)
	hasher := sha256.New()

	phoneNumbers := generatePhoneNumbers()
	lastPhoneNumber := ""
	var lastHash []byte = nil
	for phoneNumber := range phoneNumbers {
		hasher.Reset()
		_, err := hasher.Write(sharedSaltBytes)
		if err != nil {
			panic("Error: " + err.Error())
		}
		_, err = hasher.Write([]byte(phoneNumber))
		if err != nil {
			panic("Error: " + err.Error())
		}

		lastPhoneNumber = phoneNumber
		lastHash = hasher.Sum(make([]byte, hasher.BlockSize()))
	}

	fmt.Println("last number:", lastPhoneNumber)
	fmt.Printf("last hash (%d bytes): %s", len(lastHash), hex.EncodeToString(lastHash))
}
