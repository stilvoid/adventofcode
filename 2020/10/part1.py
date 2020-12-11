data = [int(x) for x in open("input").read().strip().split("\n")]
data.append(0)
print(data)

def find(path):
    last = path[-1]

    options = [
        path + [x]
        for x in data
        if x not in path
        if last-3 <= x <= last-1
    ]

    if len(options) == 0:
        yield path

    for option in options:
        if len([x for x in data if x not in option and x > option[-1]]) > 0:
            continue

        if option[-1] == 0:
            yield option
        else:
            for f in find(option):
                yield f

paths = [[x] for x in data if x == max(data)]

for path in paths:
    for option in find(path):
        if len(option) == len(data):
            diffs = {3: 1}
            for i in range(len(option)-1):
                d = option[i] - option[i+1]
                if d not in diffs:
                    diffs[d] = 0
                diffs[d] += 1
            print(diffs)
            print(diffs[3] * diffs[1])
            break
