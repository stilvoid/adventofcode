allergens = {}
ingredients = {}

for line in open("input").read().strip().split("\n"):
    ing, alg = line.split(" (contains ")

    ing = ing.split(" ")
    alg = alg[:-1].split(", ")

    for i in ing:
        if i not in ingredients:
            ingredients[i] = 0

        ingredients[i] += 1

        for a in alg:
            if a not in allergens:
                allergens[a] = {}

            if i not in allergens[a]:
                allergens[a][i] = 0

            allergens[a][i] += 1

results = {}

while any([
    len(alg.keys()) > 1
    for alg in allergens.values()
]):

    for key in results.keys():
        if key in allergens:
            del allergens[key]

    for a, alg in allergens.items():
        if a in results:
            continue

        choices = [
            ing
            for ing, count in alg.items()
            if count == max(alg.values())
        ]

        if len(choices) == 1:
            ing = choices[0]

            results[a] = ing

            for blg in allergens.values():
                if ing in blg:
                    del blg[ing]

print(sum([
    count
    for ing, count
    in ingredients.items()
    if ing not in results.values()
]))
