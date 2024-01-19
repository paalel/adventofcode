import json


with open("goals.txt", "r") as f:
    goals: list[int] = list(map(int, f.readline().rstrip().split(",")))

# split pattern [int, float], [int, float], [int, float] in a single line
with open("bets.txt", "r") as f:
    bets: list[tuple[int, float]] = json.loads(f.readline().rstrip())


initial = 50000
value = initial

for goal, t in zip(goals, bets):
    goal_guess, odds = t
    bet_value = round(value * 0.175)
    value += round(bet_value * odds) if goal >= goal_guess else -bet_value


print(initial - value)
