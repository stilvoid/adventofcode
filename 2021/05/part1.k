input: 0:"input"  / Read from file
input: {x@0 2}'" "\'input
input: ","\''input
input: 0$input

/ Filter to only straight lines
input: input@&{[row]|/({row[0][x]=row[1][x]}'!2)}'input

max_x: |/,/{[row]{x[0]}'row}'input
max_y: |/,/{[row]{x[1]}'row}'input

/ Plot
grid: ((max_x+1),(max_y+1))#0

plot: {[row]
    min_x: &/{x[0]}'row;
    max_x: |/{x[0]}'row;
    min_y: &/{x[1]}'row;
    max_y: |/{x[1]}'row;
    xs: (min_x)_!max_x+1;
    ys: (min_y)_!max_y+1;
    xs: ((#xs)|(#ys))#xs;
    ys: ((#xs)|(#ys))#ys;
    :{xs[x],ys[x]}'!#xs
    }

coords: ,/plot'input

{grid[x[1];x[0]]+: 1}'coords

+/,/1<grid

\\
