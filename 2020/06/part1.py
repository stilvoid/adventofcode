count = 0

for group in open("input").read().split("\n\n"):
    answers = {
        c: True
        for c in group
        if c not in (" ", "\n")
    }

    count += len(answers.keys())

print(count)
