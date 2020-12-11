data = {int(x): {"parts":[]} for x in open("input").read().strip().split("\n")}
data[0] = {}

for x in data.keys():
    data[x]["parts"] =[
        y for y in data
        if x-3 <= y <= x-1
    ]

print(data)

def count(x):
    if x == 0:
        return 1

    if "count" not in data[x]:
        data[x]["count"] = sum(count(y) for y in data[x]["parts"])

    return data[x]["count"]

print(max(data.keys()))
print(count(max(data.keys())))
