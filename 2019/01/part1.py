import math

with open("input") as f:
    data = f.read()

fuel = 0

for module in data.split("\n"):
    if module != "":
        fuel = fuel + (int(module) // 3) - 2

print(fuel)
