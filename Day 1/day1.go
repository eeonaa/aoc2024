package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("inputs/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	output := string(file)
	output = strings.ReplaceAll(output, "   ", "\n")
	output = strings.ReplaceAll(output, "\r", "")
	result := strings.Split(output, "\n")

	firstCollumn := []int{}
	secondCollumn := []int{}

	for i := 0; i < len(result); i++ {
		child, err := strconv.Atoi(result[i])
		if err != nil {
			fmt.Println(err)
		}

		if i%2 == 0 {
			firstCollumn = append(firstCollumn, child)
		} else {
			secondCollumn = append(secondCollumn, child)
		}
	}

	part1Total := part1(firstCollumn, secondCollumn)
	part2Total := part2(firstCollumn, secondCollumn)

	fmt.Println("Total Distance:", part1Total)
	fmt.Println("Total Similarity:", part2Total)
}

func part1(first []int, second []int) int {
	sort.Ints(first[:])
	sort.Ints(second[:])

	var abs float64
	var total int
	for i := 0; i < len(first); i++ {
		abs = math.Abs(float64(first[i] - second[i]))
		total = total + int(abs)
	}
	return total
}

func part2(first []int, second []int) int {
	var total int
	var amount int
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				amount++
			}
		}
		total = total + (first[i] * amount)
		amount = 0
	}
	return total
}
