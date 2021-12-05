input: 0:"input"  / Read from file
input: {x@0 2}'" "\'input
input: ","\''input
input: 0$input

/ Filter to only straight lines
straight: {[row]|/({row[0;x]=row[1;x]}'!2)}
diagonal: {[row](_abs(row[0;0]-row[1;0]))=(_abs(row[0;1]-row[1;1]))}

input: input@&{[row] straight[row]|diagonal[row]}'input

max_x: |/,/{[row]{x[0]}'row}'input
max_y: |/,/{[row]{x[1]}'row}'input

/ Plot
grid: ((max_x+1),(max_y+1))#0

range: {
    min: x&y;
    max: x|y;
    nums: (min)_!max+1;
    ::[x>y;|nums;nums]
}

plot: {[row]
    xs: range[row[0;0];row[1;0]];
    ys: range[row[0;1];row[1;1]];
    xs: ((#xs)|(#ys))#xs;
    ys: ((#xs)|(#ys))#ys;
    :{xs[x],ys[x]}'!#xs
    }

coords: ,/plot'input

{grid[x[1];x[0]]+: 1}'coords

+/,/1<grid

\\
