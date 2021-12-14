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

# Convert start string to pairs

chain = {}

def insert(chain, pair, n=1):
    if pair not in chain:
        chain[pair] = 0

    chain[pair] += n

for i in range(len(start)-1):
    pair = start[i] + start[i+1]
    insert(chain, pair)

counts = {}
for c in start:
    insert(counts, c)

def step(chain, counts):
    new_chain = {}

    for pair, new in rules.items():
        if pair in chain:
            c = chain[pair]

            left = pair[0] + new
            right = new + pair[1]
            
            chain[pair] -= c
            insert(new_chain, left, c)
            insert(new_chain, right, c)
            insert(counts, new, c)

    for pair, value in chain.items():
        if value > 0:
            insert(new_chain, pair, value)

    return new_chain

for i in range(40):
    chain=step(chain, counts)

print(max(counts.values()) - min(counts.values()))
