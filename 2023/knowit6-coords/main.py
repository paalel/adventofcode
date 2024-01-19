from math import sqrt

with open("rute.txt", "r") as f:
    rute = [line.strip().split(",") for line in f]


pos = rute[0]
dist = 0

for route in rute[1:]:
    dist += sqrt(
        (int(route[0]) - int(pos[0])) ** 2 + (int(route[1]) - int(pos[1])) ** 2
    )
    pos = (route[0], route[1])
print(dist // 1000 * 9)
