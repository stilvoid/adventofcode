input: 0:"input"  / Read from file
input: ,/0$'","\'input

min: &/input
max: |/input
ns: (min)_!max

&/{+/{[n] +/'{!1+_abs(x-n)}'input}[x]}'ns

\\
