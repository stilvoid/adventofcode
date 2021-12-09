input: 0:"input"  / Read from file
input: 0$''input

t: ~input=9

/ Pad edges
t: {0,x,0}'t
t: +{0,x,0}'+t

/ Save a copy
m: t

/ Number each part uniquely
t*: (^t)#!*/^t

/ Find max number of each horiz/vert neighbour
iter: {m*|/(x;1!x;-1!x;1!'x;-1!'x)}

/ Repeat til done
t: iter/t

/ Remove zeros
t: ,/t
t: t@&~t=0

/ Find the size of each group
s: {#x}'(=,/t)

/ Multiply the top three sizes
*/3#s@>s

\\
