package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./program <hex>")
		return
	}
	hexSTRING := os.Args[1]
	hexRawByte, err := hex.DecodeString(hexSTRING)
	if err != nil {
		fmt.Println(err)
		return
	}
	chunks := binaryToBase64Chunks(bytesToBinary(hexRawByte))
	var output string
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for _, chunk := range chunks {
		value, _ := strconv.ParseInt(chunk, 2, 64)
		output += string(base64Table[value])
	}
	fmt.Println(output)
	return
}

func bytesToBinary(bytes []byte) string {
	var binaryString strings.Builder
	for _, b := range bytes {
		binaryString.WriteString(fmt.Sprintf("%08b", b))
	}
	return binaryString.String()
}

func binaryToBase64Chunks(binaryString string) []string {
	var chunks []string
	for i := 0; i < len(binaryString); i += 6 {
		end := i + 6
		if end > len(binaryString) {
			end = len(binaryString)
		}
		chunks = append(chunks, binaryString[i:end])
	}
	return chunks
}
