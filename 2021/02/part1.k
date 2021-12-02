input: 0:"input"  / Read from file
input: {" "\x}'input  / Split rows on strings
input: {(x[0];0$x[1])}'input  / Convert second item of each row to int

match: {y[0]~x}
find: {y@&match[x]'y}
sum: {+/{x@1}'find[x;y]}

h: sum["forward";input]
v: sum["down";input]
v-: sum["up";input]

h*v

\\
