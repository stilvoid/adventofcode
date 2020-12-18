import numpy

data = open("input").read().strip().split("\n")

def reduce(line):
    parts = []
    part = ""
    level = 0
    for i, a in enumerate(line):
        part += a

        if a == "(":
            level += 1
            if level == 1:
                parts.append(part[:-1])
                part = part[-1]
        elif a == ")":
            level -= 1
            if level == 0:
                parts.append(str(eval(part[1:-1])))
                part = ""

    parts.append(part)

    return "".join(parts)

def eval(line):
    return numpy.prod([
        sum([int(n) for n in part.split(" + ")])
        for part in reduce(line).split(" * ")
    ])

print(sum(eval(line) for line in data))
