import math

orderrules = {}
updates = []

with open("input-day5.txt") as input_file:
    for line in input_file:

        if '|' in line:
            x, xi = line.strip().split("|")

            if not x in orderrules:
                rules = []
                rules.append(xi)
                orderrules[x] = rules
            else:
                orderrules[x].append(xi)
            continue

        if line == '\n':
            continue

        update = []
        update += line.strip().split(",")

        updates.append(update)
        
def check_update(update) -> tuple[bool, int, int]:
    for i in range(1, len(update)):

        if not update[i] in orderrules.keys():
            continue

        for j in range(0,i):
            if update[j] in orderrules[update[i]]:
                # no bueno
                print("no bueno \n{}\n {} is before {}\n".format(update, update[j], update[i]))
                return False, i, j
            
    print("bueno \n{}\n".format(update))
    return True, 0, 0
    
def middle_nr(update) -> int:
    return int(update[int(math.ceil(len(update)/2))-1])

sum = 0
resum = 0
for update in updates:
    good, i, j = check_update(update)
    if good:
        sum += middle_nr(update)
    else:
        while not good:
            update[i], update[j] = update[j], update[i]
            good, i, j = check_update(update)
        resum += middle_nr(update)

print(sum)
print(resum)