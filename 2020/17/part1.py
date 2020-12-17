cubes = {
    (x, y, 0): c
    for y, line in enumerate(open("input").read().strip().split("\n"))
    for x, c in enumerate(line)
}

def set_range():
    global mins, maxs

    keys = cubes.keys()

    xs = {key[0] for key in keys}
    ys = {key[1] for key in keys}
    zs = {key[2] for key in keys}

    mins=(min(xs), min(ys), min(zs))
    maxs=(max(xs)+1, max(ys)+1, max(zs)+1)

def get(x, y, z):
    return cubes.get((x, y, z), ".")

def output():
    for z in range(mins[2], maxs[2]):
        print(f"z={z}")
        for y in range(mins[1], maxs[1]):
            print("".join(
                get(x, y, z)
                for x in range(mins[0], maxs[0])
            ))
        print()
    print()

def active_neighbours(x, y, z):
    return len([
        True
        for iz in range(z-1, z+2)
        for iy in range(y-1, y+2)
        for ix in range(x-1, x+2)
        if iz != z or iy != y or ix != x
        if cubes.get((ix, iy, iz)) == "#"
    ])

def cycle():
    global cubes, zr, yr, xr

    set_range()

    result = {}

    for z in range(mins[2]-1, maxs[2]+1):
        for y in range(mins[1]-1, maxs[1]+1):
            for x in range(mins[0]-1, maxs[0]+1):
                n = active_neighbours(x, y, z)
                c = cubes.get((x, y, z), ".")

                if c == "#" and n not in (2, 3):
                    c = "."
                elif c == "." and n == 3:
                    c = "#"

                result[(x, y, z)] = c

    cubes = result

    set_range()

for i in range(6):
    cycle()

print(len([c for c in cubes.values() if c == "#"]))
