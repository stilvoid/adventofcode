start, buses = open("input").read().strip().split("\n")
start = int(start)
buses = [int(i) for i in buses.split(",") if i != "x"]

chosen = None
earliest = None

for bus in buses:
    e = (start // bus + 1) * bus

    if earliest is None or e < earliest:
        earliest = e
        chosen = bus

print(chosen * (earliest - start))
