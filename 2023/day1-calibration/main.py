import re

mapping = {
    "zero": "0",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}
numbers = "|".join(mapping.keys())


def get_first(line, part2):
    """
    Get the first number in the line
    """
    regex = r"\d"
    if part2:
        regex = regex + "|" + numbers

    match = re.search(regex, line)
    if not match:
        raise ValueError("No number found in line")

    g = match.group(0)
    return mapping[g] if g in mapping else g


def get_last(line, part2):
    """
    Get the last number in the line using a negative lookahead.
    """
    regex = r"\d"
    if part2:
        regex = regex + "|" + numbers

    regex = rf"({regex})?.*({regex}).*$"
    match = re.search(regex, line)
    if not match:
        raise ValueError("No number found in line")

    g = match.groups()[-1]
    return mapping[g] if g in mapping else g


def solve(file, part2: bool = False):
    total = 0

    with open(file, "r") as f:
        lines = f.readlines()

    for line in lines:
        a = get_first(line, part2=part2)
        b = get_last(line, part2=part2)
        total += int(a + b)

    return total


if __name__ == "__main__":
    test1 = solve("test_input.txt", part2=False)
    assert test1 == 142
    print("Test 1 passed!")
    print("Solution 1:", solve("input.txt", part2=False))

    test2 = solve("test_input2.txt", part2=True)
    assert test2 == 281
    print("Test  2 passed!")
    print("Solution 2:", solve("input.txt", part2=True))
