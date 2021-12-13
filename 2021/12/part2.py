with open("input") as f:
    data = [line.strip().split("-") for line in f]

data += [line[::-1] for line in data]

def can_visit(opt, visited):
    if opt not in visited:
        return True

    if any([count>1 for count in visited.values()]):
        return False

    if opt in ("start", "end"):
        return False

    return visited[opt]<2

def route(path):
    last = path[-1]

    if last == "end":
        return [path]

    opts = [line[1] for line in data if line[0] == last]

    visited = [part for part in path if part.islower()]
    visited = {
        part: len([p for p in visited if p == part])
        for part in set(visited)
    }

    opts = [opt for opt in opts if can_visit(opt, visited)]

    if len(opts) == 0:
        return [path]

    results = [route(path+[opt]) for opt in opts]

    return [r for rs in results for r in rs]

routes = route(["start"])

routes = [r for r in routes if r[-1] == "end"]

for r in routes:
    print(r)
print(len(routes))
