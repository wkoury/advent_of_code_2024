package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("input.txt")
	var left []int
	var right []int

	for _, line := range lines {
		split := strings.Split(line, " ")
		left = append(left, mustInt(split[0]))
		right = append(right, mustInt(split[3]))
	}

	// sort left & right by descending values
	sort.Slice(left, func(i, j int) bool { return left[i] > left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] > right[j] })

	fmt.Println(similarityScore(left, right))
}

func similarityScore(left []int, right []int) int {
	ret := 0
	for _, v := range left {
		ret += count(v, right) * v
	}

	return ret
}

// given a list of numbers, count the number of times a number is in the given list
func count(q int, right []int) int {
	count := 0
	for _, v := range right {
		if v == q {
			count++
		}
	}

	return count
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func mustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
	}
	contents := string(bytes)

	return strings.Split(contents, "\n")
}
