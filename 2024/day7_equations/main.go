package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result   uint64
	elements []uint64
}

type InputData []Equation

func (i *InputData) Append(row string) error {

	fields := strings.Fields(row)

	result, err := strconv.ParseUint(fields[0][:len(fields[0])-1], 10, 64)
	if err != nil {
		return err
	}
	elements := []uint64{}

	for i := 1; i < len(fields); i++ {
		num, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			return err
		}
		elements = append(elements, num)
	}

	*i = append(*i, Equation{result: result, elements: elements})
	return nil
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
		err := data.Append(line)
		if err != nil {
			return data, err
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return data, nil
}

func findAllSolutions1(elements []uint64) []uint64 {
	if len(elements) == 1 {
		return []uint64{elements[0]}
	}

	solutions := []uint64{}
	a, b := elements[0], elements[1]

	addtion := []uint64{a + b}
	addtion = append(addtion, elements[2:]...)
	solutions = append(solutions, findAllSolutions1(addtion)...)

	multiplication := []uint64{a * b}
	multiplication = append(multiplication, elements[2:]...)
	solutions = append(solutions, findAllSolutions1(multiplication)...)

	return solutions

}

func solve1(data InputData) (uint64, error) {

	var score uint64 = 0
	for _, eq := range data {
		solutions := findAllSolutions1(eq.elements)
		if slices.Contains(solutions, eq.result) {
			score += eq.result
		}
	}

	return score, nil
}

func findAllSolutions2(elements []uint64) []uint64 {
	if len(elements) == 1 {
		return []uint64{elements[0]}
	}

	solutions := []uint64{}
	a, b := elements[0], elements[1]

	addtion := []uint64{a + b}
	addtion = append(addtion, elements[2:]...)
	solutions = append(solutions, findAllSolutions2(addtion)...)

	multiplication := []uint64{a * b}
	multiplication = append(multiplication, elements[2:]...)
	solutions = append(solutions, findAllSolutions2(multiplication)...)

	concatResult, err := strconv.ParseUint(strconv.FormatUint(a, 10)+strconv.FormatUint(b, 10), 10, 64)

	if err != nil {
		panic(err)
	}

	concat := []uint64{concatResult}
	concat = append(concat, elements[2:]...)
	solutions = append(solutions, findAllSolutions2(concat)...)

	return solutions

}

func solve2(data InputData) (uint64, error) {

	var score uint64 = 0
	for _, eq := range data {
		solutions := findAllSolutions2(eq.elements)
		if slices.Contains(solutions, eq.result) {
			score += eq.result
		}
	}

	return score, nil
}

func main() {

	testData, err := readInput("test_input.txt")
	if err != nil {
		fmt.Println("Error when reading test input:", err)
		return
	}

	var expectedTestResult1 uint64 = 3749
	var expectedTestResult2 uint64 = 11387

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
