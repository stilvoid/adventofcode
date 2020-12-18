data = open("input").read().strip().split("\n")

def get(line):
    left, right = "", ""

    if line[0] == "(":
        level = 1
        i = 1
        while level > 0:
            if line[i] == "(":
                level += 1
            elif line[i] == ")":
                level -= 1
            
            left = left + line[i]

            i += 1

        left = eval(left[:-1])
        right = line[i:].strip()
    elif " " not in line:
        return line, ""
    else:
        left, right = line.split(" ", 1)

    return left, right

def eval(line):
    left, right = get(line)

    left = int(left)

    if right == "":
        return left

    op, right = right.split(" ", 1)

    right, remaining = get(right)
    right = int(right)

    if op == "+":
        left = left + right
    elif op == "*":
        left = left * right

    if remaining == "":
        return left

    return eval(f"{left} {remaining}")

print(sum(eval(line) for line in data))
