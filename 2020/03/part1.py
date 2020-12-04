world = [line.strip() for line in open("input")]
x, y = 0, 0
r, d = 3, 1

count = 0

while True:
    x += r
    y += d

    if y >= len(world):
        break

    if world[y][x % len(world[y])] == "#":
        count += 1

print(count)
