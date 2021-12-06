input: 0:"input"  / Read from file
input: ,/0$'","\'input

freqs: 9#0
{{freqs[x]+:1}'x}[input]
freqs

tick: {[f] f[7]+: f[0]; 1!f}

out: 256 tick/freqs
+/out
