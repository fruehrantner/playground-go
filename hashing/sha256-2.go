package main

import (
		"fmt"
		"crypto/sha256"
)

func main() {
	data := []byte("abcde")
	fmt.Printf("%x\n", sha256.Sum256(data))
}
