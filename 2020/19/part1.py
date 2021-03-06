import re

rules_input, data = open("input").read().strip().split("\n\n")

char_re = re.compile(r"^\"(.+)\"$")

rules = {}
for line in rules_input.split("\n"):
    name, rule = line.split(": ")
    rules[name] = {
        "line": rule,
    }

def parse(n):
    if "parsed" in rules[n]:
        return rules[n]["parsed"]

    line = rules[n]["line"]

    m = char_re.match(line)
    if m is not None:
        rules[n]["parsed"] = m.group(1)

    else:
        rules[n]["parsed"] = "(?:" + "|".join([
            "".join([parse(part) for part in parts.split(" ")])
            for parts in line.split(" | ")
        ]) + ")"

    return rules[n]["parsed"]

parse("0")

final_re = re.compile(f'^{rules["0"]["parsed"]}$')

print(len([line for line in data.split("\n") if final_re.match(line)]))
