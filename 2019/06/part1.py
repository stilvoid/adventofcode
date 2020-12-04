refs = {}
orbs = {}

with open("input") as f:
    data = f.read()

# Construct the orbit data
for row in data.strip().split("\n"):
    left, right = row.split(")")

    if left not in refs:
        refs[left] = {}

    if right not in refs:
        refs[right] = {}

    refs[left][right] = refs[right]

def walk(d, depth):
    yield depth * len(d)

    for child in d.values():
        for v in walk(child, depth+1):
            yield v

print(sum(walk(refs["COM"], 1)))
