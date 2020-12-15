start = [0, 13, 1, 8, 6, 15]
counts = {}
times = {}

for i in range(len(start) - 1):
    n = start[i]

    if n not in counts:
        counts[n] = 0

    counts[n] += 1
    times[n] = i

last = start[-1]
for i in range (len(start), 2020):
    if last not in counts:
        nxt = 0
    else:
        nxt = i - times[last] - 1

    if last not in counts:
        counts[last] = 0

    counts[last] += 1
    times[last] = i - 1

    last = nxt

    print(nxt)
