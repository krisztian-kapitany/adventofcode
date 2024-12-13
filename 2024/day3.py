import re
pattern = re.compile("mul\((\d+),(\d+)\)")

sum = 0
with open("input-day3.txt") as input_file:
    for line in input_file:
        for match in re.finditer(pattern, line):
            sum += int(match.group(1)) * int(match.group(2))

print(sum)


dopattern = re.compile("do\(\)(.*?)don't\(\)")

singleline = "do()"
with open("input-day3.txt") as input_file:
    for line in input_file:
        singleline += line.strip()
singleline += "don't()"

sum2 = 0
for match in re.finditer(dopattern, singleline):
    for innermatch in re.finditer(pattern, match.group()):
            sum2 += int(innermatch.group(1)) * int(innermatch.group(2))

print(sum2)
