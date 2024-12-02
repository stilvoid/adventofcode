with open("data") as f:
    nums = [line.strip().split("   ") for line in f]

lefts, rights = zip(*nums)

lefts = [int(l) for l in lefts]
rights = [int(r) for r in rights]

nums = zip(sorted(lefts), sorted(rights))

print(sum([abs(n[0] - n[1]) for n in nums]))
