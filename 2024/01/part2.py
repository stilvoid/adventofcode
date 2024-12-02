import math

with open("data") as f:
    nums = [line.strip().split("   ") for line in f]

lefts, rights = zip(*nums)

lefts = [int(l) for l in lefts]
rights = [int(r) for r in rights]

sims = [sum([1 for r in rights if r == l]) for l in lefts]

print(sum([math.prod(a) for a in zip(lefts, sims)]))
