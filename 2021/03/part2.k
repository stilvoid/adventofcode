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

calc: {[f;t]
    n:: -1;
    e:: ();
    out: {1<#x}iter[f]/t;
    bin2dec[out[0]]}

calc[epsilon;input] * calc[gamma;input]

\\
