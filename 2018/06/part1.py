data = open("input").read().strip()

coords = [
    (int(left), int(right))
    for left, right
    in [
        row.split(", ")
        for row
        in data.split("\n")
    ]
]

# Find limits of x and y
maxx, maxy = 0, 0
for coord in coords:
    if coord[0] > maxx:
        maxx = coord[0]

    if coord[1] > maxy:
        maxy = coord[1]

# Find the closest point to a coord
def closest_point(x, y):
    min_dist = maxx * maxy
    closest = None

    for i, coord in enumerate(coords):
        dist = abs(coord[0] - x) + abs(coord[1] - y)

        if dist == min_dist:
            closest = None
        elif dist < min_dist:
            min_dist = dist
            closest = i

    return closest

# Now count up the winners
counts = {}
remove = set()
for y in range(maxy + 1):
    line = ""

    for x in range(maxx + 2):
        p = closest_point(x, y)

        if p is None:
            line += "."
        else:
            if x == 0 or y == 0 or x == maxx or y == maxy:
                remove.add(p)

            if x == coords[p][0] and y == coords[p][1]:
                line += chr(ord("A") + p)
            else:
                line += chr(ord("a") + p)

        if p is not None:
            if p not in counts:
                counts[p] = 0

            counts[p] += 1

    print(line)

# Print out the winner
keys = [key for key in counts.keys() if key not in remove]

winner = sorted(keys, key=lambda k: counts[k])[-1]

print(counts[winner])
