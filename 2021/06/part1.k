input: 0:"input"  / Read from file
input: ,/0$'","\'input

tick: {:[x=0; 6,8; x-1]}

iter: {,/tick'x}

output: 80 iter/input
#output

\\
