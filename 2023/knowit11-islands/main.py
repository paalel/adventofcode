def remove_connected(arr: list[list[str]], i: int, j: int) -> list[list[str]]:
    arr[i][j] = "."
    n = len(arr)
    m = len(arr[0])

    # Check all 8 direction
    if i < n - 1 and arr[i + 1][j] == "X":
        arr = remove_connected(arr, i + 1, j)
    if 0 < i and arr[i - 1][j] == "X":
        arr = remove_connected(arr, i - 1, j)
    if j < m - 1 and arr[i][j + 1] == "X":
        arr = remove_connected(arr, i, j + 1)
    if 0 < j and arr[i][j - 1] == "X":
        arr = remove_connected(arr, i, j - 1)

    if i < n - 1 and j < m - 1 and arr[i + 1][j + 1] == "X":
        arr = remove_connected(arr, i + 1, j + 1)
    if 0 < i and 0 < j and arr[i - 1][j - 1] == "X":
        arr = remove_connected(arr, i - 1, j - 1)
    if i < n - 1 and 0 < j and arr[i + 1][j - 1] == "X":
        arr = remove_connected(arr, i + 1, j - 1)
    if 0 < i and j < m - 1 and arr[i - 1][j + 1] == "X":
        arr = remove_connected(arr, i - 1, j + 1)

    return arr


def solve(filename: str) -> int:
    with open(filename, "r") as f:
        arr = [[c for c in line.strip()] for line in f.readlines()]

    n = len(arr)
    m = len(arr[0])
    counter = 0

    for i in range(0, n):
        for j in range(0, m):
            if arr[i][j] == "X":
                arr = remove_connected(arr, i, j)
                counter += 1

    return counter


print("test 1")
assert solve("test_input.txt") == 1
print("test 2")
assert solve("test_input2.txt") == 2
print("test 3")
assert solve("test_input3.txt") == 2
print("test 4")
assert solve("test_input4.txt") == 3


print("All tests passed")
print(solve("kart.txt"))
