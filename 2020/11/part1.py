import copy

data = [
    [c for c in line]
    for line in open("input").read().strip().split("\n")
]

def get_adjacent(x, y):
    return [
        data[j][i]
        for j in range(y-1, y+2)
        for i in range(x-1, x+2)
        if (j != y or i != x)
        and 0 <= j < len(data)
        and 0 <= i < len(data[j])
    ]

def output():
    for line in data:
        print("".join(line))
    print()

while True:
    changes = {}
    
    for y, line in enumerate(data):
        for x, c in enumerate(line):
            if c == "L":
                if all(s != "#" for s in get_adjacent(x, y)):
                    changes[(x, y)] = "#"
            elif c == "#":
                if len([s for s in get_adjacent(x, y) if s == "#"]) >= 4:
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
