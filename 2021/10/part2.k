input: 0:"input"

openers: "([{<"
closers: ")]}>"

score: {:[x~")"; 1; x~"]"; 2; x~"}"; 3; x~">"; 4; 0]}

check: {[char; c]
    ops: char=openers;
    :[@c; c; |/ops; (closers@&ops),c; 0=#c; char; c[0]~char; (1_)c; char]}

iter: {[line]
    c: ();
    n: 0;
    do[#line; c: check[line[n];c]; n: n+1];
    :[@c; 0; {(x*5)+y}/score'c]}

scores: iter'input
scores: scores@&~0=scores
scores: scores@<scores
scores@(#scores)%2
\\
