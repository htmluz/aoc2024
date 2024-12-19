package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("erro lendo arquivo", err)
	}
	defer file.Close()
	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])
		left = append(left, n1)
		right = append(right, n2)
	}
	sort.Ints(left)
	sort.Ints(right)

	result := 0
	for i := range left {
		if left[i] < right[i] {
			result += right[i] - left[i]	
		} else {
			result += left[i] - right[i]	
		}
	}
	fmt.Println(result)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("erro lendo arquivo", err)
	}
	defer file.Close()
	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])
		left = append(left, n1)
		right = append(right, n2)
	}
	sort.Ints(left)
	sort.Ints(right)

	allcount := 0
	for _, val := range left {
		similarity := 0
		idx := sort.Search(len(right), func(i int) bool {
			return right[i] >= val
		})
		if idx < len(right) && val == right[idx] {
			count := 0 
			for i := idx; i < len(right) && right[i] == val; i++ {
				count++
			}
			similarity += val * count
		}
		allcount += similarity
	}
	fmt.Println(allcount)
}