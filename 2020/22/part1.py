decks = open("input").read().strip().split("\n\n")

decks = [
    [int(n) for n in deck.split("\n")[1:]]
    for deck in decks
]

r = 0
while all([len(deck) > 0 for deck in decks]):
    r += 1
    print(f"-- Round {r} --")

    for deck in decks:
        print(deck)

    winner = [
        i
        for i, deck in enumerate(decks)
        if deck[0] == max([deck[0] for deck in decks])
    ]

    if len(winner) != 1:
        print("BROKEN!")
    winner = winner[0]

    decks[winner].extend([decks[winner][0]] + [deck[0] for i, deck in enumerate(decks) if i != winner])

    decks = [
        deck[1:]
        for deck in decks
    ]

    print(winner)

print("-- END -- ")
print(winner)
for deck in decks:
    print(deck)

print(sum([
    n * (len(decks[winner]) - i)
    for i, n in enumerate(decks[winner])
]))
