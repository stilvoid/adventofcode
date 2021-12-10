input: 0:"input"

openers: "([{<"
closers: ")]}>"

score: {:[x~")"; 3; x~"]"; 57; x~"}"; 1197; x~">"; 25137; 0]}

check: {[char; c]
    ops: char=openers;
    :[@c; c; |/ops; (closers@&ops),c; 0=#c; char; c[0]~char; (1_)c; char]}

iter: {[line]
    c: ();
    n: 0;
    do[#line; c: check[line[n];c]; n: n+1];
    :score[c]}

+/iter'input

\\
