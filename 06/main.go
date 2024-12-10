package main

import (
	"bufio"
	"fmt"
	"os"
)

var position int = 1

func main() {
	readData()

	fmt.Println(position)
}

func readData() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(len(fileScanner.Text()))
		fmt.Println(fileScanner.Text())
	}
}
