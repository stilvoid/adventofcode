import itertools

data = [int(line.strip()) for line in open("input")]

for c in itertools.combinations(data, 2):
    if c[0] + c[1] == 2020:
        print(c[0] * c[1])
        break
