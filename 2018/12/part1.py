data = open("input").read().strip().split("\n")

initial = list(data[0].split(": ")[1])

data = data[2:]

state_map = {}

for line in data:
    in_state, out_state = line.split(" => ")

    state_map[in_state] = out_state

state = ["."] * len(initial)
state += initial
state += ["."] * len(initial)

newstate = ["."] * len(state)

offset = len(initial)

print("    " + " " * len(initial) + "0")

for i in range(20):
    for n in range(len(state) - 2):
        sample = state[n:n+5]

        newstate[n+2] = state_map.get("".join(sample), ".")

    state = newstate
    newstate = ["."] * len(state)

    print("{:2d}: {}".format(i, "".join(state)))

print(sum([n - offset if state[n] == "#" else 0 for n in range(len(state))]))
