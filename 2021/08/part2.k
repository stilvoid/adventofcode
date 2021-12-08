input: 0:"input"  / Read from file
signals: {10#x}'" "\'{x@0}'"|"\'input
outputs: {-4#x}'" "\'{x@1}'"|"\'input

signals: {signals[x],outputs[x]}'!#signals

permute :{:[1<x;,/(>:'(x,x)#1,x#0)[;0,'1+_f x-1];,!x]}
perms: {"abcdefg"@x}'permute[7]

possible:((1 1 1 0 1 1 1)
    (0 0 1 0 0 1 0)
    (1 0 1 1 1 0 1)
    (1 0 1 1 0 1 1)
    (0 1 1 1 0 1 0)
    (1 1 0 1 0 1 1)
    (1 1 0 1 1 1 1)
    (1 0 1 0 0 1 0)
    (1 1 1 1 1 1 1)
    (1 1 1 1 0 1 1))

check: {[opt;answer] |/{(|/{x=answer}'opt)~x}'possible}
checkall: {[opts;answer] &/{check[x][answer]}'opts}
solve: {[opts] (perms@&checkall[opts]'perms)@0}

convert: {[a;s] {x@0}[&{({x _in s}'a)~x}'possible]}
convert2: {[a;ss] 0$,/$convert[a]'ss}

iter: {[n]
    sigs: signals[n];
    outs: outputs[n];
    answer: solve[sigs];

    out: convert2[answer;outs];

    `0:$out
    `0:"\n"

    :out
}

+/iter'!#signals

\\
