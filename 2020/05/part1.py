code = {
    "F": "0",
    "B": "1",
    "L": "0",
    "R": "1",
}

highest = 0

for line in open("input"):
    line = line.strip()

    line = "".join([code[x] for x in line])

    row, column = int(line[:7], 2), int(line[7:], 2)

    seat = row * 8 + column

    if seat > highest:
        highest = seat

print(highest)
