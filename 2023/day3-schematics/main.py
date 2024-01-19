from dataclasses import dataclass
import re


@dataclass
class Number:
    value: int
    row: int
    start: int
    end: int

    @property
    def range(self) -> range:
        return range(self.start, self.end)


@dataclass
class Gear:
    row: int
    col: int


def parse(filename: str) -> tuple[list[Number], list[Gear], list[str]]:
    numbers: list[Number] = []
    gears: list[Gear] = []
    matrix: list[str] = []
    with open(filename, "r") as f:
        for index, line in enumerate(f.readlines()):
            matrix.append(line.strip())
            for match in re.finditer(r"\d+", line):
                numbers.append(
                    Number(
                        value=int(match.group()),
                        row=index,
                        start=match.start(),
                        end=match.end(),
                    )
                )

            for match in re.finditer(r"\*", line):
                gears.append(
                    Gear(
                        row=index,
                        col=match.start(),
                    )
                )

    return numbers, gears, matrix


def get_neighbors(matrix: list[str], row: int, start: int, end: int) -> list[str]:
    neighbors: list[str] = []
    width = len(matrix[0])
    for i in range(max(0, row - 1), min(row + 2, width)):
        neighbors.append(matrix[i][max(0, start - 1) : min(end + 1, width)])
    return neighbors


def solve1(filename):
    numbers, _, matrix = parse(filename)

    total = 0
    for number in numbers:
        if any(
            re.search(r"[^0-9.]", neighbor)
            for neighbor in get_neighbors(matrix, number.row, number.start, number.end)
        ):
            total += number.value

    return total


def solve2(filename):
    numbers, gears, matrix = parse(filename)
    total = 0

    for gear in gears:
        neighbors = []
        for i in range(max(0, gear.row - 1), min(gear.row + 2, len(matrix[0]))):
            for j in range(max(0, gear.col - 1), min(gear.col + 2, len(matrix[0]))):
                for candidate in [n for n in numbers if n.row == i and j in n.range]:
                    if candidate.value not in neighbors:
                        neighbors.append(candidate.value)

        if len(neighbors) == 2:
            total += neighbors[0] * neighbors[1]

    return total


if __name__ == "__main__":
    test1 = solve1("test_input.txt")
    assert test1 == 4361
    print("Test 1 passed!")
    print("Part 1:", solve1("input.txt"))

    test2 = solve2("test_input.txt")
    assert test2 == 467835
    print("Test 2 passed!")
    print("Part 2:", solve2("input.txt"))
