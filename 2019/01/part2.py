import math

with open("input") as f:
    data = f.read()

fuel = 0

for module in data.split("\n"):
    if module != "":
        f = (int(module) // 3) - 2

        while f > 0:
            fuel = fuel + f
            f = (f // 3) - 2

print(fuel)
