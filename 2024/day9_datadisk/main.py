test_input = "2333133121414131402"
expected_result_task1 = 1928

data = ""
with open("input.txt", "r") as file:
    data += file.readline().strip()


def read_memory_string(memory_string):
    decompressed_memory = []
    space = False
    id = 0

    for b in memory_string:
        if space:
            decompressed_memory.extend(["."] * int(b))
        else:
            decompressed_memory.extend([str(id)] * int(b))
            id += 1

        space = not space

    return decompressed_memory


def compress_memory_layout(mem):
    def last_file_index():
        for i in range(len(mem) - 1, -1, -1):
            if mem[i] != ".":
                yield i
        yield -1

    def first_free_memory_index():
        for i in range(len(mem) - 1):
            if mem[i] == ".":
                yield i
        yield -1

    last_file_index_gen = last_file_index()
    first_free_memory_index_gen = first_free_memory_index()

    while True:
        empty_index = next(first_free_memory_index_gen)
        file_index = next(last_file_index_gen)

        if empty_index == -1 or file_index == -1:
            raise ValueError("empty adress not found")

        if empty_index < file_index:
            mem[empty_index] = mem[file_index]
            mem[file_index] = "."

        elif file_index < empty_index:
            break

    return mem


def solve(memory_string):
    memory_layout = read_memory_string(memory_string)
    compressed_memory_layout = compress_memory_layout(memory_layout)

    res = 0
    for index, id in enumerate(compressed_memory_layout):
        if id == ".":
            break
        res += index * int(id)

    return res


test_result1 = solve(test_input)
assert (
    expected_result_task1 == test_result1
), f"Test failed: {expected_result_task1} != {test_result1}"

print(f"Solution to task 1: {solve(data)}")
