package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("/tmp/private_key.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Here is my private key: %s\n", string(b))
}
