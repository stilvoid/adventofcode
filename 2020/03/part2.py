import math

world = [line.strip() for line in open("input")]
slopes = ((1, 1), (3, 1), (5, 1), (7, 1), (1, 2))

answers = []

for (r, d) in slopes:
    x, y = 0, 0
    count = 0

    while True:
        x += r
        y += d

        if y >= len(world):
            break

        if world[y][x % len(world[y])] == "#":
            count += 1

    answers.append(count)

print(math.prod(answers))
