import itertools, math

print([
    math.prod(c)
    for c in itertools.combinations((int(line) for line in open("input").readlines()), 3)
    if sum(c) == 2020
])
