package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type InputData [][]int

func (i *InputData) Append(row string) {
	fields := strings.Fields(row)
	rowData := make([]int, len(fields))
	for i, f := range fields {
		num, _ := strconv.Atoi(f)
		rowData[i] = num
	}
	*i = append(*i, rowData)
}

func readInput(filename string) (InputData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := InputData{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data.Append(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return data, nil
}

func isSafe(row []int) bool {
	isIncreasing := row[0] < row[1]
	isDecreasing := row[0] > row[1]
	if !isIncreasing && !isDecreasing {
		return false
	}

	for j := 1; j < len(row); j++ {
		diff := row[j-1] - row[j]
		if (diff > -1 || diff < -3) && isIncreasing {
			return false
		}
		if (diff < 1 || diff > 3) && isDecreasing {
			return false
		}
	}
	return true
}

func solve1(data InputData) (int, error) {
	numSafe := 0

	for i := 0; i < len(data); i++ {
		if isSafe(data[i]) {
			numSafe++
		}
	}
	return numSafe, nil
}

func solve2(data InputData) (int, error) {
	numSafe := 0
	for i := 0; i < len(data); i++ {
		row := data[i]
		if isSafe(row) {
			numSafe++
			continue
		}
		for j := 0; j < len(row); j++ {
			rowWithoutIndex := slices.Delete(slices.Clone(row), j, j+1)
			if isSafe(rowWithoutIndex) {
				numSafe++
				break
			}
		}
	}
	return numSafe, nil
}

func main() {

	testData := InputData{{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}

	expectedTestResult1 := 2
	expectedTestResult2 := 4

	testSolution1, testErr := solve1(testData)

	if testErr != nil || testSolution1 != expectedTestResult1 {
		fmt.Println("Wrong solution or error in test 1:", expectedTestResult1, testSolution1, testErr)
		return
	}

	data, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error when reading input:", err)
		return
	}
	solution1, err := solve1(data)
	if err != err {
		fmt.Println("Error when solving problem 1:", err)
		return
	}
	fmt.Println("Solution 1: ", solution1)

	testSolution2, testErr := solve2(testData)
	if testErr != nil || testSolution2 != expectedTestResult2 {
		fmt.Println("Wrong solution or error in test 2:", expectedTestResult2, testSolution2, testErr)
		return
	}

	solution2, err := solve2(data)
	if err != err {
		fmt.Println("Error when solving problem 2:", err)
		return
	}
	fmt.Println("Solution 2: ", solution2)

}
