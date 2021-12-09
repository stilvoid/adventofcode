input: 0:"input"  / Read from file
input: 0$''input

testn: {[row;n] (row[n-1]>row[n])&(row[n]<row[n+1])}
testrow: {[row] testn[10,row,10;1+!#row]}

result: (testrow'input)&(+testrow'+input)

+/1+,/{input[x]@&result[x]}'!#result

\\
