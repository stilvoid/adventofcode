prog = open("input").read().strip().split("\n")
mem = {}

mask = {}

for line in prog:
    left, right = line.split(" = ")

    if left == "mask":
        mask = {
            35 - i: int(n)
            for i, n in enumerate(right)
            if n != "X"
        }

        continue

    index = int(left[4:-1])
    value = int(right)

    for i, m in mask.items():
        if m == 1:
            value |= 1 << i
        else:
            value &= ~(1 << i)

    mem[index] = value

print(sum(mem.values()))
