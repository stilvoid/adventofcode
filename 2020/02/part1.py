import re

valid = 0

for line in open("input").readlines():
    policy, password = line.strip().split(": ")
    times, char = policy.split(" ")
    mn, mx = times.split("-")

    count = len([c for c in password if c == char])

    if int(mn) <= count <= int(mx):
        valid += 1

print(valid)
