count = 0

for group in open("input").read().split("\n\n"):
    group = group.strip()

    answers = {}

    for c in group:
        if c not in (" ", "\n"):
            if c not in answers:
                answers[c] = 0
            answers[c] += 1

    people = len(group.split("\n"))

    counted = [
        key
        for key, value in answers.items()
        if value == people
    ]

    count += len(counted)

print(count)
