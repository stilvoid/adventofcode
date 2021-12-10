input: 0:"input"

openers: "([{<"
closers: ")]}>"

score: {:[x~")"; 3; x~"]"; 57; x~"}"; 1197; x~">"; 25137; 0]}

check: {[c;char]
    ops: char=openers;
    :[@c; c; |/ops; (closers@&ops),c; 0=#c; char; c[0]~char; (1_)c; char]}

results: {check/[();x]}'input
results: results@&{@x}'results
+/score'results
\\
