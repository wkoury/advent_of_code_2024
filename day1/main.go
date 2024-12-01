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
		fmt.Println(split)
		left = append(left, mustInt(split[0]))
		right = append(right, mustInt(split[3]))
	}

	fmt.Println(left)
	fmt.Println(right)

	if len(left) != len(right) {
		panic("lens not equal")
	}

	// sort left & right by descending values
	sort.Slice(left, func(i, j int) bool { return left[i] > left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] > right[j] })

	fmt.Println(left)
	fmt.Println(right)

	distanceSum := 0

	for ii := 0; ii < len(left); ii++ {
		// get the absolute value of the difference
		fmt.Printf("left[%d]: %d\nright[%d]: %d\n", ii, left[ii], ii, right[ii])
		distanceSum += absInt(left[ii] - right[ii])
	}

	fmt.Println(distanceSum)
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
