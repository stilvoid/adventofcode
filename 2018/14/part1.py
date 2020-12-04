import sys

recipes = "37"
elves = [0, 1]

target=18

while len(recipes) < target + 10:
    new = str(sum(int(recipes[pos]) for pos in elves))

    recipes += new
    
    for i, pos in enumerate(elves):
        elves[i] = (elves[i] + 1 + int(recipes[pos])) % len(recipes)

print(recipes[target:target+10])
