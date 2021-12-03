input: 0:"input"  / Read from file
input: 0$''input  / Convert to 2d array of ints

epsilon: {[a]{~x<(#a)%2.0}'+/a}
gamma: {~epsilon[x]}

bin2dec: {[a]+/{x[y]*2^y}[|a]'!#a}

`0:"input:\n"
input

iter: {[f;t]
    e:: f[t];
    n:: n+1;
    t[&{x[n]=e[n]}'t]}

n: -1
e: ()
ox: {1<#x}iter[epsilon]/input
ox: bin2dec[ox[0]]
ox

n: -1
e: ()
co: {1<#x}iter[gamma]/input
co: bin2dec[co[0]]
co

ox * co

\\
