string = ""
count = 0

for i in range(100000):
    if str(i) not in string:
        string += str(i)
        count += 1

print(count)
