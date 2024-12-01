package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var firstColumn []int
	var secondColumn []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integers:", line)
			continue
		}

		firstColumn = append(firstColumn, a)
		secondColumn = append(secondColumn, b)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return firstColumn, secondColumn, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve1(arr1, arr2 []int) (int, error) {
	slices.Sort(arr1)
	slices.Sort(arr2)

	length1 := len(arr1)
	length2 := len(arr2)
	if length1 != length2 {
		return 0, errors.New("Input data not same length in solve1")
	}

	similarity := 0
	for i := 0; i < length1; i++ {
		similarity += abs(arr1[i] - arr2[i])
	}
	return similarity, nil
}

func solve2(arr1, arr2 []int) (int, error) {
	hm := make(map[int]int)

	for _, v := range arr2 {
		_, ok := hm[v]
		if ok {
			hm[v]++
		} else {
			hm[v] = 1
		}
	}

	similarity := 0
	for _, v := range arr1 {
		numApps, ok := hm[v]
		if ok {
			similarity += v * numApps
		}
	}

	return similarity, nil

}

func main() {
	testData1 := []int{3, 4, 2, 1, 3, 3}
	testData2 := []int{4, 3, 5, 3, 9, 3}

	testSolution1, testErr := solve1(testData1, testData2)
	if testErr != nil || testSolution1 != 11 {
		fmt.Println("Wrong solution (not 11) or error", testSolution1, testErr)
		return
	}

	data1, data2, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error when reading input", err)
		return
	}
	solution1, err := solve1(data1, data2)
	if err != err {
		fmt.Println("Error when solving problem 1", err)
		return
	}
	fmt.Println("Solution 1: ", solution1)

	testSolution2, testErr := solve2(testData1, testData2)
	if testErr != nil || testSolution2 != 31 {
		fmt.Println("Wrong solution (not 31) or error", testSolution2, testErr)
		return
	}

	solution2, err := solve2(data1, data2)
	if err != err {
		fmt.Println("Error when solving problem 2", err)
		return
	}
	fmt.Println("Solution 2: ", solution2)

}
