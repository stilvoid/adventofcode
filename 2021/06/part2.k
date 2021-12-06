input: 0:"input"  / Read from file
input: ,/0$'","\'input

freqs: +/'{(input=x)}'!9

tick: {[f] f[7]+: f[0]; 1!f}

out: 256 tick/freqs
+/out

\\
