package main

import (
	"fmt"
	"os"
)

func main() {
	readWriteToken := os.Getenv("ACCESS_TOKEN")
	fmt.Printf("Here is my secret access token: %s\n", readWriteToken)
}
