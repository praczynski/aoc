package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// number := utils.ReadFileToString("input.txt")
	number := "2333133121414131402"
	var sectors []string
	counter := 0

	// Parse the disk map into sectors
	for g, v := range strings.Split(number, "") {
		ile, _ := strconv.Atoi(v)
		if g%2 == 0 {
			// Add file blocks with the current counter as ID
			for i := 0; i < ile; i++ {
				sectors = append(sectors, strconv.Itoa(counter))
			}
			counter++
		} else {
			// Add free space blocks
			for i := 0; i < ile; i++ {
				sectors = append(sectors, ".")
			}
		}
	}

	length := len(sectors)
	buffer := []string{}
	spaceBuffer := []int{}
	var dumped bool
	dupa := 0

	// Compact the disk by moving file blocks to the leftmost free space
	for j := length - 1; j > dupa; j-- {
		if sectors[j] != "." {
			if len(buffer) == 0 || dumped || sectors[j] == buffer[0] {
				buffer = append(buffer, sectors[j])
				sectors[j] = "."
				continue
			}
		}
		dumped = false
		if len(buffer) > 0 {
			for i := dupa; i < j; i++ {
				if sectors[i] == "." {
					spaceBuffer = append(spaceBuffer, i)
					if len(spaceBuffer) >= len(buffer) {
						break
					}
				} else {
					spaceBuffer = []int{}
				}
			}
			if len(spaceBuffer) >= len(buffer) {
				for g := 0; g < len(buffer); g++ {
					sectors[spaceBuffer[g]] = buffer[g]
				}
				dupa = len(buffer)
				dumped = true
				spaceBuffer = []int{}
			}
			j++
		}
		buffer = []string{}
	}

	fmt.Println(sectors)

	// Calculate the checksum
	result := 0
	for i, v := range sectors {
		if v == "." {
			continue
		}
		value, _ := strconv.Atoi(v)
		result += i * value
	}
	fmt.Println(result)
}
