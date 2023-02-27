package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}
