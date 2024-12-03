package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InputData string

func (i *InputData) Append(row string) {
	*i += InputData(row)
}

func readInput(filename string) (InputData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data InputData = ""
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
	type Multiplication struct {
		x int
		y int
	}

	multiplications := []Multiplication{}

	instructionRegex := regexp.MustCompile(`mul\(([[0-9]{1,3},[0-9]{1,3})\)`)
	for _, match := range instructionRegex.FindAllStringSubmatch(string(data), -1) {
		nums := strings.Split(match[1], ",")
		x, errX := strconv.Atoi(nums[0])
		y, errY := strconv.Atoi(nums[1])

		if errX != nil || errY != nil {
			fmt.Println(errX, errY)
			return 0, errY
		}
		multiplications = append(multiplications, Multiplication{x: x, y: y})
	}
	product := 0
	for _, m := range multiplications {
		product += m.x * m.y
	}

	return product, nil
}

func solve2(data InputData) (int, error) {
	cleanRegex := regexp.MustCompile(`don't\(\).*?do\(\)`)
	return solve1(InputData(cleanRegex.ReplaceAllString(string(data), "")))
}

func main() {

	testData1 := InputData("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	testData2 := InputData("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	expectedTestResult1 := 161
	expectedTestResult2 := 48

	testSolution1, testErr := solve1(testData1)

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

	testSolution2, testErr := solve2(testData2)
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
