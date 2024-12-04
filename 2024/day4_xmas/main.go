package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputData [][]rune

func (i *InputData) Append(row string) {
	rowData := make([]rune, len(row))
	for i, v := range row {
		rowData[i] = v
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

func check(data *InputData, searchString string, row int, col int, dx int, dy int) bool {
	height, width := len(*data), len((*data)[0])

	for index, searchRune := range []rune(searchString) {
		x, y := row+index*dx, col+index*dy

		if x < 0 || x >= height || y < 0 || y >= width {
			return false
		}
		if searchRune != (*data)[x][y] {
			return false
		}
	}

	return true
}

func solve1(data InputData) (int, error) {

	searchDirections := [][]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}
	numXmas := 0
	height, width := len(data), len((data)[0])
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			for _, dir := range searchDirections {
				dx, dy := dir[0], dir[1]
				if check(&data, "XMAS", row, col, dx, dy) {
					numXmas += 1
				}
			}
		}
	}

	return numXmas, nil
}

func solve2(data InputData) (int, error) {
	height, width := len(data), len((data)[0])
	numXhyphenMas := 0
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			if !check(&data, "MAS", row-1, col-1, 1, 1) &&
				!check(&data, "SAM", row-1, col-1, 1, 1) {
				continue
			}
			if !check(&data, "MAS", row+1, col-1, -1, 1) &&
				!check(&data, "SAM", row+1, col-1, -1, 1) {
				continue
			}

			numXhyphenMas++
		}
	}

	return numXhyphenMas, nil
}

func main() {

	testData := InputData{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	expectedTestResult1 := 18
	expectedTestResult2 := 9

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
