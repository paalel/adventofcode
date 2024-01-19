def parse(filename: str) -> list[tuple[list[int], list[int]]]:
    with open(filename, "r") as f:
        lines = f.readlines()

    cards = [line.rstrip().split(": ")[1] for line in lines]
    split_cards = [card.split(" | ") for card in cards]
    return [
        (
            [int(num) for num in card[0].split(" ") if num != ""],
            [int(num) for num in card[1].split(" ") if num != ""],
        )
        for card in split_cards
    ]


def solve1(filename: str) -> int:
    total = 0
    cards = parse(filename)
    for card in cards:
        winners, actual = card
        winning_numbers = [a for a in actual if a in winners]
        num_matches = len(winning_numbers)

        total += pow(2, (num_matches - 1)) if num_matches > 0 else 0

    return total


if __name__ == "__main__":
    test1 = solve1("test_input.txt")
    assert test1 == 13
    print("Test 1 successful")
    print("Solution 1:", solve1("input.txt"))
