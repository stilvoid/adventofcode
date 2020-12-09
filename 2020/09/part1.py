import itertools

l = 25

data = [int(x) for x in open("input").read().strip().split("\n")]

for i in range(l, len(data)):
    d = data[i]

    valid = {
        a + b
        for a, b in itertools.permutations(data[i-l:i], 2)
    }

    if d not in valid:
        print(d)
        break
