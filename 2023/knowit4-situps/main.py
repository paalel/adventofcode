def calc(reps: list[int]):
    curr_val = reps[0]
    curr_len = 1

    max_val = curr_val
    max_len = 1

    for idx, rep in enumerate(reps[:-1]):
        if rep <= reps[idx + 1]:
            curr_len += 1
            curr_val += reps[idx + 1]
        else:
            curr_len = 1
            curr_val = reps[idx + 1]

        if curr_len > max_len or (curr_len == max_len and curr_val > max_val):
            max_val = curr_val
            max_len = curr_len

    return max_val


# fmt: off
test = [133, 266, 174, 295, 228, 257, 75, 41, 370, 125, 188, 284, 301, 219, 276, 134, 315, 190, 183, 381, 12, 351, 384, 151, 255, 231, 232, 205, 95, 0, 97]
# fmt: on

assert calc(test) == 898, f"Expected 898, got {calc(test)}"

with open("reps.txt", "r") as f:
    reps = list(map(int, f.readline().split(", ")))

print(calc(reps))
