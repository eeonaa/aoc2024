package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
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
	output = strings.ReplaceAll(output, "\r", "")
	split := strings.Split(output, "\n")

	var result [][]int
	for i := 0; i < len(split); i++ {
		temp := strings.Split(split[i], " ")
		var tempIntArr []int
		for j := 0; j < len(temp); j++ {
			tempInt, err := strconv.Atoi(temp[j])
			if err != nil {
				fmt.Println(err)
			}

			tempIntArr = append(tempIntArr, tempInt)
			if len(tempIntArr) == len(temp) {
				result = append(result, tempIntArr)
			}
		}
	}

	total := checkInOrder(result)

	fmt.Println("Total Safe Damped Reports:", total)
}

func checkInOrder(result [][]int) int {
	var total int
	for i := 0; i < len(result); i++ {
		var tempAscending []int
		var tempDescending []int

		tempAscending = append(tempAscending, result[i]...)
		sort.Ints(tempAscending[:])

		tempDescending = reverseInts(tempAscending)

		if reflect.DeepEqual(result[i], tempAscending) || reflect.DeepEqual(result[i], tempDescending) {
			if secondCheck(result[i]) == 1 {
				total += secondCheck(result[i])
			} else {
				total += checkOutOfOrder(result[i])
			}
		} else {
			total += checkOutOfOrder(result[i])
		}
	}
	return total
}

func checkOutOfOrder(result []int) int {
	var total int
	var completed bool = false
	var tempResult []int

	tempResult = append(tempResult, result...)

	for i := 0; i < len(result); i++ {
		if completed {
			continue
		}

		tempResult = deleteElement(tempResult, i)

		var tempAscending []int
		tempAscending = append(tempAscending, tempResult...)

		var tempDescending []int
		tempDescending = append(tempDescending, tempResult...)

		sort.Ints(tempAscending[:])
		sort.Ints(tempDescending[:])
		tempDescending = reverseInts(tempDescending)

		if reflect.DeepEqual(tempResult, tempAscending) || reflect.DeepEqual(tempResult, tempDescending) {
			tempTotal := total
			total += secondCheck(tempResult)
			if tempTotal != total {
				completed = true
			}
		}

		tempResult = nil
		tempResult = append(tempResult, result...)
	}

	return total
}

func secondCheck(result []int) int {
	var total int
	var sumArr []int
	var tempSumArr []int
	for j := 0; j < len(result); j++ {
		var sum float64
		if !(j == 0) {
			sum = math.Abs(float64(result[j-1]) - float64(result[j]))
			sumArr = append(sumArr, int(sum))
		}
	}

	for j := 0; j < len(sumArr); j++ {
		if sumArr[j] <= 3 && sumArr[j] >= 1 {
			tempSumArr = append(tempSumArr, sumArr[j])
		}
	}
	if len(tempSumArr) == len(result)-1 {
		total++
	}

	return total
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
