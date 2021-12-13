with open("input") as f:
    data = [line.strip().split("-") for line in f]

data += [line[::-1] for line in data]

def route(path):
    last = path[-1]

    if last == "end":
        return [path]

    opts = [line[1] for line in data if line[0] == last]

    visited = [part for part in path if part.islower()]

    opts = [opt for opt in opts if opt not in visited]

    if len(opts) == 0:
        return [path]

    results = [route(path+[opt]) for opt in opts]

    return [r for rs in results for r in rs]

routes = route(["start"])

routes = [r for r in routes if r[-1] == "end"]

print(len(routes))
