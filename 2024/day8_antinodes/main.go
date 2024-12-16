package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputData [][]rune

func (i *InputData) Append(row string) {
	*i = append(*i, []rune(row))
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

type Node struct {
	x   int
	y   int
	dir int
}

func solve(data InputData, harmonize bool) (int, error) {

	towerMap := make(map[rune][][2]int)
	for i, arr := range data {
		for j, r := range arr {
			if r == '.' {
				continue
			}
			_, ok := towerMap[r]
			if ok {
				towerMap[r] = append(towerMap[r], [2]int{i, j})
			} else {
				towerMap[r] = [][2]int{{i, j}}
			}
		}
	}

	alreadySeenNodes := make(map[[2]int]bool)
	calc_num_antinodes := func(a, b [2]int) int {

		score := 0
		diff := [2]int{a[0] - b[0], a[1] - b[1]}
		nodes := []Node{{x: a[0], y: a[1], dir: 1}, {x: b[0], y: b[1], dir: -1}}

		for _, node := range nodes {
			resonance := 0
			if harmonize {
				resonance = -1
			}

		resonanceLoop:
			for {
				resonance++
				x, y := node.x+node.dir*resonance*diff[0], node.y+node.dir*resonance*diff[1]
				antinode := [2]int{x, y}

				if alreadySeenNodes[antinode] && harmonize {
					continue resonanceLoop
				}

				if alreadySeenNodes[antinode] && !harmonize {
					break resonanceLoop
				}

				if 0 > x || x > len(data)-1 || 0 > y || y > len(data[0])-1 {
					break resonanceLoop
				}

				alreadySeenNodes[antinode] = true
				score++

				if !harmonize {
					break resonanceLoop
				}
			}
		}

		return score
	}

	score := 0
	for _, towers := range towerMap {
		for i := 0; i < len(towers); i++ {
			for j := i + 1; j < len(towers); j++ {
				a := towers[i]
				b := towers[j]
				score += calc_num_antinodes(a, b)
			}
		}
	}

	return score, nil
}

func solve1(data InputData) (int, error) {
	return solve(data, false)
}

func solve2(data InputData) (int, error) {
	return solve(data, true)
}

func main() {

	expectedTestResult1 := 14
	expectedTestResult2 := 34

	testData, err := readInput("test_input.txt")
	if err != nil {
		fmt.Println("Error when reading test input:", err)
		return
	}

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
