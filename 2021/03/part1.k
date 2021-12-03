input: 0:"input"  / Read from file
input: 0$''input  / Convert to 2d array of ints

epsilon: {x>(#input)%2}'+/input

gamma: ~epsilon

bin2dec: {[a]+/{x[y]*2^y}[|a]'!#a}

epsilon: bin2dec[epsilon]
gamma: bin2dec[gamma]

epsilon * gamma

\\
