from dataclasses import dataclass
from typing import Dict


@dataclass
class Round:
    red: int
    green: int
    blue: int

    def __init__(self, line: str):
        split_array = [x.split(" ") for x in line.split(", ")]
        colors: Dict[str, int] = {color: int(amount) for amount, color in split_array}
        self.red = colors.get("red", 0)
        self.green = colors.get("green", 0)
        self.blue = colors.get("blue", 0)


def parse_file(filename: str) -> list[list[Round]]:
    game = []

    with open(filename) as f:
        lines = f.readlines()

    for line in lines:
        rounds = line.rstrip().split(": ")[1].split("; ")
        parsed_round = []
        for round in rounds:
            parsed_round.append(Round(round))

        game.append(parsed_round)

    return game


def solve1(filename):
    total = 0
    game = parse_file(filename)
    for index, rounds in enumerate(game):
        if all(
            round.red <= 12 and round.green <= 13 and round.blue <= 14
            for round in rounds
        ):
            total += index + 1

    return total


def solve2(filename):
    total = 0
    game = parse_file(filename)

    for rounds in game:
        max_red = max(round.red for round in rounds)
        max_green = max(round.green for round in rounds)
        max_blue = max(round.blue for round in rounds)
        total += max_red * max_green * max_blue

    return total


if __name__ == "__main__":
    test1 = solve1("test_input.txt")
    assert test1 == 8
    print("Solution to part 1 is:", solve1("input.txt"))

    test2 = solve2("test_input.txt")
    print(test2)
    assert test2 == 2286
    print("Solution to part 2 is:", solve2("input.txt"))
