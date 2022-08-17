package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func scan(sendChan chan string){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "error"){
			sendChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil{
			log.Default().Fatalf("error scanning line: %s", scanner.Text())
		}
	}
}

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	recieve := make(chan string)
	go func(){
		scan(recieve)
	}()
	for resp := range recieve{
		fmt.Println(resp)
	}


	


	// TODO: Look for lines in the STDIN reader that contain "error" and output them.
}
