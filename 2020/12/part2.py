sx, sy = 0, 0
wx, wy = 10, 1

dirs = ("N", "E", "S", "W")

for line in open("input").read().strip().split("\n"):
    c, n = line[0], int(line[1:])

    if c == "N":
        wy += n
    elif c == "E":
        wx += n
    elif c == "S":
        wy -= n
    elif c == "W":
        wx -= n
    elif c == "L":
        for i in range(n//90):
            wx, wy = -wy, wx
    elif c == "R":
        for i in range(n//90):
            wx, wy = wy, -wx
    elif c == "F":
        sx += wx * n
        sy += wy * n

print(abs(sx) + abs(sy))
