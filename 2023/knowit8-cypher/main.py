import json


def encode_char(char, keys: list[int]):
    if char in [" ", "\n", ","]:
        return char

    char = {"æ": "{", "å": "}", "ø": "|"}.get(char, char)
    v = chr(ord("a") + keys[ord(char) - ord("a")])
    return {"{": "æ", "}": "å", "|": "ø"}.get(v, v)


def create_decode_map(key: list[int]):
    """
    Creates a map from encoded to decoded characters
    """

    return {encode_char(char, key): char for char in "abcdefghijklmnopqrstuvwxyzæøå"}


# Test that example works
# fmt: off
testkey= [15, 23, 21, 5, 11, 26, 8, 20, 6, 28, 2, 7, 19, 3, 22, 14, 24, 1, 18, 13, 4, 12, 27, 0, 17, 9, 25, 16, 10]
# fmt: on

test_decode_map = create_decode_map(testkey)
assert "".join(test_decode_map[char] for char in "åeh") == "jul"


# Solve todays problem
with open("cypher.txt", "r") as file:
    lines = file.readlines()

with open("keys.txt", "r") as file:
    maplist = [create_decode_map(json.loads(key)) for key in file]

count = 0
for line in lines:
    output = ""
    for word in line.rstrip().split(" "):
        keymap = maplist[count]
        decoded = "".join([keymap.get(char, char) for char in word])
        count += 1
        output += decoded + " "
    print(output)
