input: 0:"input"  / Read from file

nums: 0$'","\input[0]
called: !0

flatten: {:[@x;x;,/_f'x]}

// format into arrays of numbers
cards: 2 _ input
cards: " "\'cards
cards: flatten[0$'cards]
cards: cards@&~0N=cards
cards: -1 5 5 #cards

marks: ((#cards),5,5)#0

draw: {
    called:: nums[0]
    nums:: 1 _ nums

    marks:: marks|called=cards
}

win: {(|/(&/x))|(|/(&/+x))}

score: {[n] called * (+/(+/cards[n]*~marks[n]))}
    
play: {while[(!0)~&win'marks; draw[]]}

round: {play[]; cards:: cards[&~win'marks]; marks:: marks[&~win'marks]}

/ Play until the last hand
while[1<#cards; round[]]

/ Play to finish
play[]

score[0]

\\
