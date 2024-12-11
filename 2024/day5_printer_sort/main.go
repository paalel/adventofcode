package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type InputData struct {
	Rules map[int][]int
	Books [][]int
}

func (i *InputData) Append(line string) {
	rulePattern := regexp.MustCompile(`^(\d*)\|(\d*)$`)
	if match := rulePattern.FindStringSubmatch(line); len(match) > 0 {
		key, _ := strconv.Atoi(match[1])
		value, _ := strconv.Atoi(match[2])

		_, ok := i.Rules[key]
		if ok {
			i.Rules[key] = append(i.Rules[key], value)
		} else {
			i.Rules[key] = []int{value}
		}
		return
	}

	pages := []int{}
	for _, page := range strings.Split(line, ",") {
		pageNumber, _ := strconv.Atoi(page)
		pages = append(pages, pageNumber)
	}
	i.Books = append(i.Books, pages)

}

func readInput(filename string) (InputData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return InputData{}, err
	}
	defer file.Close()

	data := InputData{Books: [][]int{}, Rules: make(map[int][]int)}
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

func checkPageRules(book []int, pageIndex int, pageRules []int) bool {
	for _, pageRule := range pageRules {
		if !slices.Contains(book, pageRule) {
			continue
		}
		if !slices.Contains(book[pageIndex:], pageRule) {
			return false
		}
	}

	return true
}

func checkBookOrdering(book []int, rules map[int][]int) bool {
	for index, page := range book {
		pageRules, ok := rules[page]
		if !ok {
			continue
		}
		if !checkPageRules(book, index, pageRules) {
			return false
		}

	}
	return true

}

func calcScore(books [][]int) (int, error) {
	pageScore := 0
	for _, book := range books {
		if len(book)%2 == 0 {
			return 0, fmt.Errorf("Book has odd number of pages: %v", book)
		}

		middleIndex := (len(book) - 1) / 2
		pageScore += book[middleIndex]
	}

	return pageScore, nil

}

func solve1(data InputData) (int, error) {
	correctBooks := [][]int{}
	for _, book := range data.Books {
		if checkBookOrdering(book, data.Rules) {
			correctBooks = append(correctBooks, book)
		}
	}

	return calcScore(correctBooks)
}

func findProblemIndex(book []int, rules map[int][]int) int {
	for index, page := range book {

		pageRules, ok := rules[page]
		if !ok {
			continue
		}
		if !checkPageRules(book, index, pageRules) {
			return index
		}

	}
	return -1
}

func reorderBook(book []int, rules map[int][]int) ([]int, error) {

	reorganizedBook := slices.Clone(book)

	loopCounter := 0
	for {
		loopCounter++
		problemIndex := findProblemIndex(reorganizedBook, rules)

		if problemIndex == 0 {
			panic("This shouldnt happen: rules only apply to later elements and every element is after the first")
		}
		if loopCounter > 10000 {
			panic("In an inf loop: probably swapping back and forth")
		}

		if problemIndex == -1 {
			break
		}

		// Assume that we can get to a correct ordering by swapping neighbours
		tmp := reorganizedBook[problemIndex]
		reorganizedBook[problemIndex] = reorganizedBook[problemIndex-1]
		reorganizedBook[problemIndex-1] = tmp
	}

	return reorganizedBook, nil
}

func solve2(data InputData) (int, error) {
	incorrectBooks := [][]int{}
	for _, book := range data.Books {
		if !checkBookOrdering(book, data.Rules) {
			incorrectBooks = append(incorrectBooks, book)
		}
	}
	reorganizedBooks := [][]int{}
	for _, book := range incorrectBooks {

		reorderedBook, err := reorderBook(book, data.Rules)
		if err != nil {
			return 0, err
		}
		reorganizedBooks = append(reorganizedBooks, reorderedBook)
	}

	return calcScore(reorganizedBooks)
}

func main() {

	testData, err := readInput("test_input.txt")
	if err != nil {
		fmt.Println("Error when reading test_input:", err)
		return
	}
	expectedTestResult1 := 143
	expectedTestResult2 := 123

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
