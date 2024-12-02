with open("data") as f:
    reports = [[int(l) for l in line.split(" ")] for line in f]

def is_safe(report):
    pairs = zip(report[:-1], report[1:])

    diffs = [n[0] - n[1] for n in pairs]

    diffs_ok = [1 <= abs(d) <= 3 for d in diffs]

    if not all(diffs_ok):
        return False

    if all(d > 0 for d in diffs):
        return True

    if all(d < 0 for d in diffs):
        return True

    return False

def dampner(report):
    return any([is_safe(report[0:n] + report[n+1:]) for n in range(len(report))])

print(len([report for report in reports if dampner(report)]))
