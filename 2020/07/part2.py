rules = {}

for line in open("input"):
    left, right = line.strip().split(" contain ")

    name = " ".join(left.split(" ")[:2])

    if name not in rules:
        rules[name] = {}

    if right[:-1] == "no other bags":
        continue

    for content in right[:-1].split(", "):
        num, a, b, _ = content.split(" ")

        content = a + " " + b

        if content not in rules[name]:
            rules[name][content] = 0

        rules[name][content] += int(num)

print(rules)

def size(name):
    return 1 + sum(value * size(key) for key, value in rules[name].items())

print(size("shiny gold"))
