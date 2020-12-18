import numpy, re

data = open("input").read().strip().split("\n")

def eval(line):
    while "(" in line:
        line = re.sub(r"\([^()]+\)", lambda m: str(eval(m.group(0)[1:-1])), line)

    return numpy.prod([
        sum([int(n) for n in part.split(" + ")])
        for part in line.split(" * ")
    ])

print(sum(eval(line) for line in data))
