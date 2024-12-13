def difference_map(levels):
    return [ levels[i+1]-levels[i] for i in range(len(levels)-1) ]

def is_safe(diffs):
    return all(-3<=x<=-1 for x in diffs) or all(1<=x<=3 for x in diffs)

def dampened_levels(levels):
    base_diffs = difference_map(levels)
    candidate = levels.copy()

    if(sum(x for x in base_diffs) < 0):
         base_diffs = [-x for x in base_diffs]

    for i,x in enumerate(base_diffs):
        if(x<1 or x>3):
            levels.pop(i)
            candidate.pop(i+1)
            break
    
    return levels, candidate

safe_count = 0
dampened_safe_count = 0
with open("input-day2.txt") as input_file:
    for ix, line in enumerate(input_file):
        l = [ int(c) for c in line.split(" ") ]

        if(is_safe(difference_map(l))):
            safe_count+=1

        else:
            c1, c2 = dampened_levels(l)
            if(is_safe(difference_map(c1)) or is_safe(difference_map(c2))):
                dampened_safe_count+=1

print(safe_count)
print(safe_count + dampened_safe_count)

