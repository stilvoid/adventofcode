a = 1
x, y = 0, 0

dirs = ("N", "E", "S", "W")

for line in open("input").read().strip().split("\n"):
    c, n = line[0], int(line[1:])

    # Turning or forward
    if c == "F":
        c = dirs[a]
    elif c == "L":
        a = (a - n//90) % 4
    elif c == "R":
        a = (a + n//90) % 4

    # Now move
    if c == "N":
        y = y + n
    elif c == "E":
        x = x + n
    elif c == "S":
        y = y - n
    elif c == "W":
        x = x - n

print(abs(x) + abs(y))
