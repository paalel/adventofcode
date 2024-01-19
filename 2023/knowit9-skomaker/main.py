import ast
from collections import defaultdict


def solve(pairs):
    neighbor_map = defaultdict(list)
    for a, b in pairs:
        neighbor_map[a].append(b)
        neighbor_map[b].append(a)

    # Find start elf - this is the elf with only one neighbor
    single_elves = [
        elf for elf, neighbors in neighbor_map.items() if len(neighbors) == 1
    ]
    elf_queue = [single_elves[0]]

    while len(elf_queue) < len(pairs) + 1:
        neigbors = neighbor_map[elf_queue[-1]]
        next = neigbors[0] if neigbors[0] not in elf_queue else neigbors[1]
        elf_queue.append(next)

    middle = len(elf_queue) // 2
    return elf_queue[middle - 1] + elf_queue[middle]


test_pairs = [(1, 2), (4, 3), (2, 3)]
assert solve(test_pairs) == 5, f"Expected 5, got {solve(test_pairs)}"


with open("rekke.txt", "r") as f:
    pairs = ast.literal_eval(f.read())

print(solve(pairs))
