refs = {}

with open("input") as f:
    data = f.read()

# Construct the orbit data
for row in data.strip().split("\n"):
    left, right = row.split(")")

    if left not in refs:
        refs[left] = None

    refs[right] = left

def path(d):
    yield d
    if refs[d] is not None:
        for p in path(refs[d]):
            yield p

you = list(path("YOU"))[::-1]
san = list(path("SAN"))[::-1]

print(you)
print(san)

for i in range(min(len(you), len(san))):
    if you[i] != san[i]:
        print(len(you) - i + len(san) - i - 2)
        break
