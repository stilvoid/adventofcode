import re

valid = 0

for line in open("input").readlines():
    policy, password = line.strip().split(": ")
    pos, char = policy.split(" ")
    a, b = pos.split("-")
    a, b = int(a), int(b)

    count = 0

    if len(password) >= a and password[a-1] == char:
        count += 1

    if len(password) >= b and password[b-1] == char:
        count += 1

    if count == 1:
        valid += 1

print(valid)
