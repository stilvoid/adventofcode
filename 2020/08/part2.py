import copy

orig = [
    {
        "op": line.split(" ")[0],
        "arg": int(line.split(" ")[1]),
    }
    for line in open("input").read().strip().split("\n")
]

def run(program):
    acc = 0
    ptr = 0

    while ptr < len(program):
        line = program[ptr]

        if "visited" in line:
            return 1, acc

        line["visited"] = True

        if line["op"] == "acc":
            acc += line["arg"]
            ptr += 1
            continue

        if line["op"] == "jmp":
            ptr += line["arg"]
            continue

        ptr += 1

    return 0, acc

for i in range(len(orig)):
    new = copy.deepcopy(orig)

    if new[i]["op"] == "nop":
        new[i]["op"] = "jmp"
    elif new[i]["op"] == "jmp":
        new[i]["op"] = "nop"

    code, acc = run(new)

    if code == 0:
        print(acc)
        break
