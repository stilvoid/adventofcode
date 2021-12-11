input: 0:"input"
input: 0$''input

pad: {+{0,x,0}'+{0,x,0}'x}
unpad: {+1_{((#x)-1)#x}[+1_((#x)-1)#x]}

shift: {+x!+y!z}

flash: {[data]
    grid: data[0];
    old: data[1];
    f: (grid>9)&~old;
    a: +/{shift[x[0];x[1];f]}'(-1 -1;0 -1;1 -1;-1 0;1 0;-1 1;0 1; 1 1);
    :(grid+a;f|old)}

iter: {[grid]
    p: pad[grid+1];
    p: (p;(^p)#0);
    q: {flash x}\p;
    q: {x@0}-1#q; / Whut? Kona is broken. "over" would have been better in the line above
    :unpad[q[0]*~q[1]]}

100 {[x] input:: iter[input]; x+(+//input=0)}/0
\\
