import sympy
import functools


@functools.lru_cache(maxsize=1000000)
def prime_cache(func):
    return func


prime = prime_cache(sympy.isprime)


def is_sum_prime(n):
    s = sum([int(i) for i in str(n)])
    if n % s != 0:
        return False

    d = n // s

    return prime(d)


def count_sum_primes(N: int) -> int:
    count = 0
    for i in range(2, N):
        if is_sum_prime(i):
            count += 1

        if i % 1000000 == 0:
            print(f"Progress: {i/N*100:.2f}%")

    return count


if __name__ == "__main__":
    print(count_sum_primes(100000000))
