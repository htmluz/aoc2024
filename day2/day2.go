package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	safe := 0
	unsafe := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		asc := false
		desc := false
		valid := true
		for x := 0; x < len(parts)-1; x++ {
			a, _ := strconv.Atoi(parts[x])
			b, _ := strconv.Atoi(parts[x + 1])
			diff := b - a
			if diff > 3 || diff == 0 || diff < -3 {
				unsafe++
				valid = false
				break
			}
			if diff > 0 {
				if desc {
					unsafe++
					valid = false
					break
				}
				asc = true
			} else if diff < 0 {
				if asc {
					unsafe++
					valid = false
					break
				}
				desc = true
			}
		}
		if valid {
			safe++
		}
	}
	fmt.Println(safe, unsafe)
}