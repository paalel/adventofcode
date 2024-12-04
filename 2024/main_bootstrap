package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type InputData // TODO

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


func solve1(data InputData) (int, error) {
  return 0, nil
}

func solve2(data InputData) (int, error) {
  return 0, nil
}

func main() {

	testData := InputData{} // TODO
	expectedTestResult1 :=
	expectedTestResult2 := 

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
