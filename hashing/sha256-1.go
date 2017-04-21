package main

import (
		"fmt"
		"io"
		"crypto/sha256"
)

func main() {
	text := "abcde"
	hash := sha256.New()
	io.WriteString(hash, text)
	fmt.Printf("%x\n", hash.Sum(nil))
}
