from typing import Generator
from sympy import sieve


def decrypt(c: str, k: int) -> str:
    if c in " .,!?":
        return c

    f = ord("A") if c.isupper() else ord("a")
    return chr((ord(c) - f - k) % 26 + f)


def num_twin_primes(n: int) -> int:
    primes = [p for p in sieve.primerange(0, n)]

    count = 0
    for i in range(len(primes) - 1):
        if primes[i + 1] - primes[i] == 2:
            count += 1

    return count


def even_binaries() -> Generator[int, None, None]:
    idx = 0
    while True:
        binary = bin(idx)[2:]
        if binary.count("1") % 2 == 0:
            yield idx

        idx += 1


if __name__ == "__main__":
    cypher = "Ojfkyezkz bvclae zisj a guomiwly qr tmuematbcqxqv sa zmcgloz."

    x = num_twin_primes(666 + pow(6, 6))
    y = even_binaries()
    solution = "".join([decrypt(c, x * next(y)) for c in cypher])

    print("Solution:", solution)
