import copy

data = [
    [c for c in line]
    for line in open("input").read().strip().split("\n")
]

def look(x, y):
    for j in range(-1, 2):
        for i in range(-1, 2):
            if j == 0 and i == 0:
                continue

            mul = 1
            while True:
                ix = x + mul * i
                jy = y + mul * j

                if jy < 0 or jy >= len(data):
                    break

                if ix < 0 or ix >= len(data[jy]):
                    break

                if data[jy][ix] != ".":
                    yield data[jy][ix]
                    break

                mul += 1

def output():
    for line in data:
        print("".join(line))
    print()

while True:
    changes = {}
    
    for y, line in enumerate(data):
        for x, c in enumerate(line):
            if c == "L":
                if not any(s == "#" for s in look(x, y)):
                    changes[(x, y)] = "#"
            elif c == "#":
                if len([l for l in look(x, y) if l == "#"]) >= 5:
                    changes[(x, y)] = "L"

    for (x, y), s in changes.items():
        data[y][x] = s

    if len(changes.keys()) == 0:
        break

print(len([
    True
    for line in data
    for c in line
    if c == "#"
]))
