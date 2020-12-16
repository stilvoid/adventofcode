rules_text, my_ticket, tickets = open("input").read().strip().split("\n\n")

rules = {}
for rule in rules_text.split("\n"):
    name, value = rule.split(": ")
    rules[name] = []

    for rng in value.split(" or "):
        rules[name].append([int(n) for n in rng.split("-")])

my_ticket = [int(n) for n in my_ticket.split("\n")[1].split(",")]

tickets = [
    [int(n) for n in ticket.split(",")]
    for ticket in tickets.split("\n")[1:]
]

def invalid_nums(ticket):
    for value in ticket:
        if not any([
            mn <= value <= mx
            for rule in rules.values()
            for mn, mx in rule
        ]):
            yield value

print(sum([
    result
    for ticket in tickets
    for result in invalid_nums(ticket)
]))
