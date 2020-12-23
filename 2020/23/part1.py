cups = [int(c) for c in open("input").read().strip()]

mn, mx = min(cups), max(cups)

pos = 0
for i in range(100):
    print(cups)

    current = cups[pos]

    sel = [
        cups[i % len(cups)]
        for i in range(pos+1, pos+4)
    ]
    cups = [cup for cup in cups if cup not in sel]

    positions = {
        cup: i
        for i, cup in enumerate(cups)
    }

    dest = current - 1

    while dest not in positions:
        dest -= 1
        if dest < mn:
            dest = mx

    print(current, sel, cups, dest)

    cups = cups[:positions[dest] + 1] + sel + cups[positions[dest]+1:]

    positions = {
        cup: i
        for i, cup in enumerate(cups)
    }

    pos = (positions[current] + 1) % len(cups)

    print()

pos = positions[1]
print("".join([
    str(cups[i % len(cups)])
    for i in range(pos + 1, pos + len(cups))
]))
