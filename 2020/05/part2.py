code = {
    "F": "0",
    "B": "1",
    "L": "0",
    "R": "1",
}

seats = {}

for line in open("input"):
    line = line.strip()

    line = "".join([code[x] for x in line])

    row, column = int(line[:7], 2), int(line[7:], 2)

    seat = row * 8 + column

    seats[seat] = True

options = {}

for seat in seats.keys():
    for x in (seat-1, seat+1):
        if x not in seats:
            if x not in options:
                options[x] = 0
            options[x] += 1

print(options)
