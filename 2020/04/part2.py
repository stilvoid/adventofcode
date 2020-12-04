import re

required = {
    "byr": lambda x: 1920 <= int(x) <= 2002,
    "iyr": lambda x: 2010 <= int(x) <= 2020,
    "eyr": lambda x: 2020 <= int(x) <= 2030,
    "hgt": lambda x: 150 <= int(x[:-2]) <= 193 if x[-2:] == "cm" else 59 <= int(x[:-2]) <= 76,
    "hcl": lambda x: re.match("^#[0-9a-f]{6}$", x) is not None,
    "ecl": lambda x: x in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
    "pid": lambda x: re.match(r"^\d{9}$", x) is not None,
}

passports = [
    {
        entry.split(":")[0]: entry.split(":")[1]
        for entry in passport.split()
    }
    for passport
    in open("input").read().split("\n\n")
]

print(len([
    passport
    for passport in passports
    if all(
        field in passport and rule(passport[field])
        for field, rule in required.items()
    )
]))
