prog = open("input").read().strip().split("\n")
mem = {}
mask = {}

def decode(m, s):
    if s == -1:
        yield []
        return

    for option in decode(m, s - 1):
        if m[s] == "X":
            yield [0] + option
            yield [1] + option
        else:
            yield [m[s]] + option


for line in prog:
    left, right = line.split(" = ")

    if left == "mask":
        mask = {
            35 - i: int(n) if n != "X" else "X"
            for i, n in enumerate(right)
        }

        continue

    left = int(left[4:-1])
    value = int(right)

    # Prepare index mask
    index = {i: n for i, n in mask.items()}
    left = bin(left)[2:]
    for i, n in enumerate(left):
        i = len(left) - i - 1

        if index[i] == 0:
            index[i] = int(n)
        

    for option in decode(index, 35):
        option = "".join([str(n) for n in option])
        mem[int(option, 2)] = value

print(sum(mem.values()))
