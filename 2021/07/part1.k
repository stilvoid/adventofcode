input: 0:"input"  / Read from file
input: ,/0$'","\'input

test: {[crabs;n] +/{_abs (x-n)}'crabs}

min: &/input
max: |/input

&/{test[input;x]}'(min)_!max

\\
