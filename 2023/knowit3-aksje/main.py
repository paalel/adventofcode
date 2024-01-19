def calc_optimal(day: list[int]) -> tuple[int, int]:
    optimal = (day[0], day[1])
    for i, v0 in enumerate(day[:-1]):
        v1 = max(day[i + 1 :])
        if v1 - v0 > optimal[1] - optimal[0]:
            optimal = (v0, v1)

    return optimal


filename = "input.txt"
money = 200000
days: list[list[int]] = []

with open(filename, "r") as file:
    for line in file:
        days.append([int(x) for x in line.split(",")])

    for day in days:
        buyprice, sellprice = calc_optimal(day)
        stocks = money // buyprice
        money_left = money - stocks * buyprice
        money = money_left + stocks * sellprice


print(money)
