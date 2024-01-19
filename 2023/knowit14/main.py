def calc(reps: list[int]):
    curr_val = reps[0]
    curr_len = 1

    max_val = curr_val
    max_len = 1

    up = True
    for idx, rep in enumerate(reps[:-1]):
        if up and rep <= reps[idx + 1]:
            curr_len += 1
            curr_val += reps[idx + 1]

        elif curr_len > 1 and rep > reps[idx + 1]:
            up = False
            curr_len += 1
            curr_val += reps[idx + 1]
        else:
            up = True
            curr_len = 1
            curr_val = reps[idx + 1] + reps[idx]

        if curr_len > max_len or (curr_len == max_len and curr_val > max_val):
            max_val = curr_val
            max_len = curr_len

    return max_val


# fmt: off
test = [16, 3, 1, 2, 9, 8, 12, 14, 19, 21, 20, 11, 2, 4, 3, 1, 7, 9]
# fmt: on

assert calc(test) == 107, f"Expected 107, got {calc(test)}"

with open("push.txt", "r") as f:
    reps = list(map(int, f.readline().split(", ")))

print(calc(reps))
