with open("log.txt", "r") as f:
    line = f.readline()

count = 0
lock = {1: False, 2: False, 3: False, 4: False, 5: False, 6: False, 7: False}
for action in line.split(", "):
    if action.startswith("klikk på "):
        num = int(action.split("klikk på ")[1])
        lock[num] = True
    elif action.startswith("klakk på "):
        num = int(action.split("klakk på ")[1])
        lock[num] = False
    else:
        print("Ukjent handling", action)
        raise ValueError("Ukjent handling")

    if all(lock.values()):
        count += 1
        lock = {1: False, 2: False, 3: False, 4: False, 5: False, 6: False, 7: False}


print(count)
