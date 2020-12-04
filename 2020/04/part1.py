required = (
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
)

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
    if all(field in passport for field in required)
]))
