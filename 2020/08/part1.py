program = [
    {
        "op": line.split(" ")[0],
        "arg": int(line.split(" ")[1]),
    }
    for line in open("input").read().strip().split("\n")
]

acc = 0
ptr = 0

while True:
    line = program[ptr]

    if "visited" in line:
        break

    line["visited"] = True

    if line["op"] == "acc":
        acc += line["arg"]
        ptr += 1
        continue

    if line["op"] == "jmp":
        ptr += line["arg"]
        continue

    ptr += 1

print(acc)
