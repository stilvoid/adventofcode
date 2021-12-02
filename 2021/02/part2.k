input: 0:"input"  / Read from file
input: {" "\x}'input  / Split rows on strings
input: {(x[0];0$x[1])}'input  / Convert second item of each row to int

input: {:[x[0]~"down";x[1],0;x[0]~"up";-x[1],0;0,x[1]]}'input  / Split into aim and movement instructions

aim: 0
depth: 0
h: 0

{aim+:x[0];h+:x[1];depth+:aim*x[1]}'input

h*depth

\\
