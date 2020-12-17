cubes = {
    0: {
        y: {
            x: c
            for x, c in enumerate(line)
        }
        for y, line in enumerate(open("input").read().strip().split("\n"))
    }
}


def set_range():
    global xr, yr, zr

    zs = cubes.keys()
    ys = {y for plane in cubes.values() for y in plane.keys()}
    xs = {x for plane in cubes.values() for line in plane.values() for x in line.keys()}

    xr = (min(xs), max(xs)+1)
    yr = (min(ys), max(ys)+1)
    zr = (min(zs), max(zs)+1)

def get(x, y, z):
    return cubes.get(z, {}).get(y, {}).get(x, ".")

def set(data, x, y, z, value):
    if z not in data:
        data[z] = {}

    if y not in data[z]:
        data[z][y] = {}

    data[z][y][x] = value

def output():
    for z in range(zr[0], zr[1]):
        print(f"z={z}")
        for y in range(yr[0], yr[1]):
            print("".join([
                get(x, y, z)
                for x in range(xr[0], xr[1])
            ]))
        print()
    print()

set_range()
output()

def active_neighbours(x, y, z):
    return len([
        True
        for iz in range(z-1, z+2)
        for iy in range(y-1, y+2)
        for ix in range(x-1, x+2)
        if iz != z or iy != y or ix != x
        if get(ix, iy, iz) == "#"
    ])

def cycle():
    global cubes, zr, yr, xr
    result = {}

    for z in range(zr[0]-1, zr[1]+1):
        for y in range(yr[0]-1, yr[1]+1):
            for x in range(xr[0]-1, xr[1]+1):
                n = active_neighbours(x, y, z)
                c = get(x, y, z)

                if c == "#" and n not in (2, 3):
                    set(result, x, y, z, ".")
                elif c == "." and n == 3:
                    set(result, x, y, z, "#")
                else:
                    set(result, x, y, z, c)

    cubes = result

    set_range()

for i in range(6):
    cycle()

print(len([c for plane in cubes.values() for line in plane.values() for c in line.values() if c == "#"]))
