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
def closest_point_dist(x, y):
    min_dist = maxx * maxy

    for coord in coords:

        if dist < min_dist:
            min_dist = dist

    return min_dist

# Now count up the winners
winners = 0
for y in range(maxy + 1):
    for x in range(maxx + 2):
        total_dist = 0
        for coord in coords:
            total_dist += abs(coord[0] - x) + abs(coord[1] - y)

        if total_dist < 10000:
            winners += 1

print(winners)
