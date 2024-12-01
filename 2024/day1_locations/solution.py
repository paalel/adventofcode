test_input = [(3, 4), (4, 3), (2, 5), (1, 3), (3, 9), (3, 3)]


def solve1(data):
    return sum(
        abs(l - r)
        for l, r in zip(sorted([x[0] for x in data]), sorted([x[1] for x in data]))
    )


def solve2(data):
    hm = {}
    for v in [x[1] for x in data]:
        hm[v] = 1 if v not in hm else hm[v] + 1

    return sum((hm[v] if v in hm else 0) * v for v in [x[0] for x in data])


with open("input.txt", "r") as file:
    assert solve1(test_input) == 11, f"Test 1 did not pass: {solve1(test_input)}"

    data = [tuple(map(int, line.split())) for line in file]
    print(f"Solution 1: {solve1(data)}")

    assert solve2(test_input) == 31, f"Test 2 did not pass: {solve2(test_input)}"
    print(f"Solution 2: {solve2(data)}")
