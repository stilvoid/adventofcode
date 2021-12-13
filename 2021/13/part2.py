dots = []
steps = []

done_dots = False
with open("input") as f:
    for line in f:
        line = line.strip()

        if line == "":
            done_dots=True
        elif not done_dots:
            dots += [tuple(int(x) for x in line.split(","))]
        else:
            line = line.split(" ")
            line = line[-1]
            line = line.split("=")
            steps += [(line[0], int(line[1]))]

print(dots)
print(steps)

max_x = 1 + max([dot[0] for dot in dots])
max_y = 1 + max([dot[1] for dot in dots])

grid = [[False for x in range(max_x)] for y in range(max_y)]

def dump(grid):
    for row in grid:
        print("".join(["#" if x else "." for x in row]))
    
    print()

for dot in dots:
    grid[dot[1]][dot[0]] = True

def foldy(grid, y):
    top, bottom = grid[0:y], grid[y+1:]
    bottom = bottom[::-1]

    for i, row in enumerate(bottom):
        for j, dot in enumerate(row):
            top[i][j] = top[i][j] or dot

    return top

def foldx(grid, x):
    left = [row[0:x] for row in grid]
    right = [row[x+1:] for row in grid]

    right = [row[::-1] for row in right]

    for i, row in enumerate(right):
        for j, dot in enumerate(row):
            left[i][j] = left[i][j] or dot

    return left

def fold(grid, step):
    if step[0] == "x":
        return foldx(grid, step[1])
    return foldy(grid, step[1])

for step in steps:
    grid = fold(grid, step)

dump(grid)
