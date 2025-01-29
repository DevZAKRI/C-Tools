package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 && len(args) != 3 {
		fmt.Println("Invalid Args")
		fmt.Println("Usage: ./main option data key(optional)")
		fmt.Println("option: enc OR dec")
		fmt.Println("data: ciphertext OR plaintext")
		fmt.Println("Key: 1-25 u can specify the key to encrypt too or decrypt from")
		fmt.Println("Key optional not added print all")
		return
	}
	option := args[0]
	data := args[1]
	if option != "enc" && option != "dec" {
		fmt.Println("Invalid Args")
		fmt.Println("Usage: ./main option data")
		fmt.Println("option: enc OR dec")
		fmt.Println("data: ciphertext OR plaintext")
		return
	}

	if len(args) == 3 {
		key := rune(args[2][0] - '0')
		var ciphertext, plaintext string
		for _, char := range data {
			switch option {
			case "enc":
				if char >= 'a' && char <= 'z' {
					ciphertext += string((char-'a'+key)%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					ciphertext += string((char-'A'+key)%26 + 'A')
				} else {
					ciphertext += string(char)
				}
			case "dec":
				if char >= 'a' && char <= 'z' {
					plaintext += string((char-'a'-key+26)%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					plaintext += string((char-'A'-key+26)%26 + 'A')
				} else {
					plaintext += string(char)
				}
			}
		}
		fmt.Println("Result:", ciphertext+plaintext)
		return
	}

	var i rune
	var ciphertext, plaintext string
	for i = 1; i <= 25; i++ {
		for _, char := range data {
			switch option {
			case "enc":
				if char >= 'a' && char <= 'z' {
					ciphertext += string((char-'a'+i)%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					ciphertext += string((char-'A'+i)%26 + 'A')
				} else {
					ciphertext += string(char)
				}
			case "dec":
				if char >= 'a' && char <= 'z' {
					plaintext += string((char-'a'-i+26)%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					plaintext += string((char-'A'-i+26)%26 + 'A')
				} else {
					plaintext += string(char)
				}
			}
		}
		if option == "enc" {
			fmt.Println("For Key:", i, "the ciphertext is:", ciphertext)
			ciphertext = ""
		} else {
			fmt.Println("For Key:", i, "the plaintext is:", plaintext)
			plaintext = ""
		}
	}
}
