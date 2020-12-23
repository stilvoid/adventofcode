start = [
    [int(n) for n in line.split("\n")[1:]]
    for line in open("input").read().strip().split("\n\n")
]

def play(decks):
    rounds = {}

    while all([len(deck) > 0 for deck in decks]):
        index = tuple(tuple(card for card in deck) for deck in decks)
        # Recursion breaker
        if index in rounds:
            return 0, decks[0]

        rounds[index] = True

        if all([
            len(deck) > deck[0]
            for deck in decks
        ]):
            winner, _ = play([[card for card in deck[1:deck[0]+1]] for deck in decks])

        else:
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

    winner = [
        i
        for i, deck in enumerate(decks)
        if len(deck) == max([len(deck) for deck in decks])
    ][0]

    return winner, decks[winner]

_, deck = play(start)
print(deck)
print(sum([
    n * (len(deck) - i)
    for i, n in enumerate(deck)
]))
