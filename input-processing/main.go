package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.
}
