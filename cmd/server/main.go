package main

import (
	"fmt"
	"os"
)

func main() {
	if err := cmd.RunServer; err != nil {
		fmt.Println(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
