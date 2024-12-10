package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	number := utils.ReadFileToString("input.txt")
	//number := "2333133121414131402"
	type Sector struct {
		id      string
		visited bool
	}
	var sectors []Sector
	counter := 0

	for g, v := range strings.Split(number, "") {
		ile, _ := strconv.Atoi(v)
		if g%2 == 0 {
			for i := 0; i < ile; i++ {
				sectors = append(sectors, Sector{strconv.Itoa(counter), false})
			}
			counter++
		} else {
			// Add free space blocks
			for i := 0; i < ile; i++ {
				sectors = append(sectors, Sector{".", false})
			}
		}
	}

	length := len(sectors)
	var buffer []string
	var spaceBuffer []int

	for j := length - 1; j >= 0; j-- {
		if sectors[j].id != "." && !sectors[j].visited {
			if len(buffer) > 0 && sectors[j].id == buffer[0] {
				buffer = append(buffer, sectors[j].id)
				sectors[j].id = "."
				continue
			}
		}
		if len(buffer) > 0 {
			for i := 0; i < j; i++ {
				if sectors[i].id == "." {
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
					sectors[spaceBuffer[g]] = Sector{buffer[g], true}
				}
				spaceBuffer = []int{}
			} else {
				for g := 0; g < len(buffer); g++ {
					sectors[j+g+1] = Sector{buffer[g], true}
				}
			}
		}

		buffer = []string{}
		if !sectors[j].visited {
			buffer = append(buffer, sectors[j].id)
			sectors[j] = Sector{".", true}
		}
	}

	result := 0
	for i, v := range sectors {
		if v.id == "." {
			continue
		}
		value, _ := strconv.Atoi(v.id)
		result += i * value
	}
	fmt.Println(result)
}
