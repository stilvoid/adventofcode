data = open("input").read().strip().split("\n")

initial = list(data[0].split(": ")[1])

data = data[2:]

state_map = {}

count = 50000000000

for line in data:
    in_state, out_state = line.split(" => ")

    state_map[in_state] = out_state

state = initial
offset = 2
last_offset = 2

print("    " + " " * offset + "0")

state = list(".." + "".join(state).strip(".") + "..")
newstate = ["."] * len(state)
print("{:2d}: {}".format(0, "".join(state)))

for i in range(50000000000):
    last_offset = offset

    # Reset
    for n in range(len(state)):
        if state[n] != ".":
            trim = n
            break

    offset -= n
    offset += 4

    state = list("...." + "".join(state).strip(".") + "....")
    newstate = ["."] * len(state)

    for n in range(len(state) - 2):
        sample = state[n:n+5]

        newstate[n+2] = state_map.get("".join(sample), ".")

    if i % 1000000 == 0:
        print(i)

    if "".join(state).strip(".") == "".join(newstate).strip("."):
        print("{:2d}: {}".format(i + 1, "".join(state)))
        print("old offset:", last_offset)
        print("new offset:", offset)
        print("DONE!")
        break

    state = newstate

print("{:2d}: {}".format(i + 1, "".join(state)))

iters = 50000000000-i-1
offset += iters * (last_offset - offset)

print(sum([n - offset if state[n] == "#" else 0 for n in range(len(state))]))
