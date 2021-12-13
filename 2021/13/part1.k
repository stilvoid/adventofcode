input: 0:"input"
split: {x@0}[&{""~x}'input]

dots: 0$'","\'split#input

steps: ((split+1)_)input
steps: {*(|x)}'" "\'steps
steps: "="\'steps

max_x: 1+(|/{*x}'dots)
max_y: 1+(|/{*(|x)}'dots)

grid: (max_y;max_x) # 0

{grid[x[1];x[0]]: 1}'dots;

print: {`0:${".#"[x]}'x}

foldy: {[grid;y]
    top: y#grid;
    bottom: |((y+1)_)grid;
    top|: bottom;
    :top}

foldx: {[grid;x]
    left: {[row] x#row}'grid;
    right: {[row] |((x+1)_)row}'grid;
    left|: right;
    :left}

iter: {[grid;step]
    :[step[0]~,"x"; foldx[grid;0$step[1]]; foldy[grid;0$step[1]]]}

g: iter[grid;steps[0]]
print[g]
+//g

\\
