package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	GUARD_UP    rune = '^'
	GUARD_DOWN  rune = '⌄'
	GUARD_LEFT  rune = '<'
	GUARD_RIGHT rune = '>'
	WALL        rune = '#'
	FLOOR       rune = '.'
	VISITED     rune = 'o'
)

var MOVEMENT = map[rune][2]int{
	GUARD_UP:    {-1, 0},
	GUARD_DOWN:  {1, 0},
	GUARD_LEFT:  {0, -1},
	GUARD_RIGHT: {0, 1},
}

var ROTATION = map[rune]rune{
	GUARD_UP:    GUARD_RIGHT,
	GUARD_RIGHT: GUARD_DOWN,
	GUARD_DOWN:  GUARD_LEFT,
	GUARD_LEFT:  GUARD_UP,
}

type Node struct {
	Typ          rune
	VisitedRight bool
	VisitedDown  bool
}

type InputData struct {
	Grid [][]Node
	X    int
	Y    int
}

func (i *InputData) get(x, y int) rune {
	node := i.Grid[x][y]
	return node.Typ
}

func (i *InputData) visit(x, y int, typ rune) {
	if typ == GUARD_RIGHT {
		i.Grid[x][y].VisitedRight = true
	}
	if typ == GUARD_DOWN {
		i.Grid[x][y].VisitedDown = true
	}
	i.Grid[x][y].Typ = VISITED

}

var (
	OUT_OF_BOUNDS = "out of bounds"
	LOOP          = "loop"
	MOVED         = "moved"
	ROTATED       = "rotated"
)

func (i *InputData) moveGuard() string {

	movement := MOVEMENT[i.get(i.X, i.Y)]
	dx, dy := i.X+movement[0], i.Y+movement[1]

	if dx < 0 || dx >= len(i.Grid) {
		i.visit(i.X, i.Y, i.Grid[i.X][i.Y].Typ)
		return OUT_OF_BOUNDS
	}

	if dy < 0 || dy >= len(i.Grid[0]) {
		i.visit(i.X, i.Y, i.Grid[i.X][i.Y].Typ)
		return OUT_OF_BOUNDS
	}

	currentNode := i.Grid[i.X][i.Y]
	nextNode := i.Grid[dx][dy]

	if currentNode.Typ == GUARD_RIGHT && currentNode.VisitedRight {
		return LOOP
	}
	if currentNode.Typ == GUARD_DOWN && currentNode.VisitedDown {
		return LOOP
	}

	if nextNode.Typ == WALL {
		i.Grid[i.X][i.Y].Typ = ROTATION[i.Grid[i.X][i.Y].Typ]
		return ROTATED
	}

	if nextNode.Typ == VISITED || nextNode.Typ == FLOOR {
		i.Grid[dx][dy].Typ = i.get(i.X, i.Y)
		i.visit(i.X, i.Y, i.get(i.X, i.Y))
		i.X, i.Y = dx, dy
		return MOVED
	}

	panic("Uhandled case in moveGuard")
}

func (i *InputData) PrintGrid() {
	for _, arr := range i.Grid {
		for _, node := range arr {
			fmt.Print(string(node.Typ))

		}
		fmt.Print("\n")
	}
}

func (i *InputData) Append(row string) {
	rowData := []Node{}

	for j, r := range row {
		if r == GUARD_UP {
			i.X, i.Y = len(i.Grid), j
		}

		rowData = append(rowData, Node{Typ: r, VisitedRight: false, VisitedDown: false})
	}

	i.Grid = append(i.Grid, rowData)
}

func (i *InputData) Clone() InputData {
	grid := make([][]Node, len(i.Grid))
	for i, row := range i.Grid {
		nodes := make([]Node, len(row))

		for j, node := range row {
			nodes[j] = Node{
				Typ:          node.Typ,
				VisitedRight: node.VisitedRight,
				VisitedDown:  node.VisitedDown,
			}
		}
		grid[i] = nodes
	}

	return InputData{
		X:    i.X,
		Y:    i.Y,
		Grid: grid,
	}
}

func readInput(filename string) (InputData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return InputData{}, err
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
	room := data.Clone()

	for {
		status := room.moveGuard()
		if status == OUT_OF_BOUNDS {
			break
		}

		if status == LOOP {
			return 0, errors.New("Encountered loop in problem 1")
		}

	}

	numVisited := 0
	for _, arr := range room.Grid {
		for _, node := range arr {
			if node.Typ == VISITED {
				numVisited++
			}
		}
	}

	return numVisited, nil
}

func solve2(data InputData) (int, error) {
	initX, initY := data.X, data.Y
	room := data.Clone()

	for room.moveGuard() != OUT_OF_BOUNDS {

	}

	visitedPositions := [][2]int{}
	for i, row := range room.Grid {
		for j, node := range row {
			if i == initX && j == initY {
				continue
			}
			if node.Typ == VISITED {
				visitedPositions = append(visitedPositions, [2]int{i, j})
			}
		}
	}

	alreadyChecked := make(map[[2]int]bool)

	numLoops := 0
	for _, pos := range visitedPositions {
		_, ok := alreadyChecked[pos]
		if ok {
			continue
		}
		alreadyChecked[pos] = true

		x, y := pos[0], pos[1]
		alteredRoom := data.Clone()
		alteredRoom.Grid[x][y].Typ = WALL

		for {
			status := alteredRoom.moveGuard()
			if status == LOOP {
				numLoops++
				break
			}
			if status == OUT_OF_BOUNDS {
				break
			}
		}
	}
	return numLoops, nil
}

func main() {

	testData, err := readInput("test_input.txt")
	if err != nil {
		fmt.Println("Error when reading test input:", err)
		return
	}
	expectedTestResult1 := 41
	expectedTestResult2 := 6

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
