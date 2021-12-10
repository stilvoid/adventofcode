input: 0:"input"

openers: "([{<"
closers: ")]}>"

score: {:[x~")"; 1; x~"]"; 2; x~"}"; 3; x~">"; 4; 0]}

check: {[c;char]
    ops: char=openers;
    :[@c; c; |/ops; (closers@&ops),c; 0=#c; char; c[0]~char; (1_)c; char]}

results: {check/[();x]}'input
results: results@&~{@x}'results
results: score''results
results: {(5*x)+y}/'results
results: results@<results
results@((#results)%2)

\\
