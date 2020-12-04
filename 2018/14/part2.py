import sys

recipes = "37"
elves = [0, 1]

target="157901"

count = 0

while True:
    count += 1

    new = str(sum(int(recipes[pos]) for pos in elves))

    recipes += new
    
    for i, pos in enumerate(elves):
        elves[i] = (elves[i] + 1 + int(recipes[pos])) % len(recipes)

    if target in recipes[-len(target)-1:]:
        break

print(recipes.find(target))
