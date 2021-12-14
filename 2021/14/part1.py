start = ""
rules = {}

done_start = False
with open("input") as f:
    i = 0
    for line in f:
        line = line.strip()

        if not done_start:
            start = line
            done_start = True
        elif line != "":
            line = line.split(" -> ")
            rules[line[0]] = line[1]

def find_all(s, sub):
    i = 0

    while i<len(s):
        i = s.find(sub, i)

        if i != -1:
            yield i
            i += 1
        else:
            i = len(s)

def step(chain):
    found = {}

    for pair, new in rules.items():
        for i in find_all(chain, pair):
            found[i] = new

    out = ""

    for i in range(len(chain)):
        out += chain[i]

        if i in found:
            out += found[i]

    return out

def count(chain):
    counts = {}

    for c in chain:
        if c not in counts:
            counts[c] = 0

        counts[c] += 1

    return counts

for i in range(10):
    start = step(start)

counts = count(start)

min_count = min(counts.values())
max_count = max(counts.values())

print(max_count - min_count)
