import numpy

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

def is_valid(ticket):
    for value in ticket:
        if not any((
            rng[0] <= value <= rng[1]
            for rule in rules.values()
            for rng in rule
        )):
            return False

    return True

num_fields = len(tickets[0])
tickets = [ticket for ticket in tickets if is_valid(ticket)] + [my_ticket]

def get_rule_location(rule):
    for i in range(num_fields):
        if all([
            any([
                rng[0] <= ticket[i] <= rng[1]
                for rng in rule
            ])
            for ticket in tickets
        ]):
            yield i

locations = {
    n: list(get_rule_location(r))
    for n, r in rules.items()
}

# Resolve down to a single option for each field
while any(len(locs) > 1 for locs in locations.values()):
    for name, locs in locations.items():
        if len(locs) == 1:
            for other in locations.keys():
                if other != name:
                    locations[other] = [
                        loc for loc in locations[other]
                        if loc != locs[0]
                    ]

locations = {name: locs[0] for name, locs in locations.items()}

print(numpy.prod([
    my_ticket[locations[name]]
    for name, rule in rules.items()
    if name.startswith("departure")
]))
