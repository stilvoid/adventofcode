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

found = set()
seeking = {"shiny gold"}
keep_going = True

while keep_going:
    for f in found:
        seeking.add(f)

    keep_going = False

    for name, rule in rules.items():
        for target in seeking:
            if target in rule:
                if name not in found:
                    found.add(name)
                    keep_going = True

print(found)

print(len(found))
