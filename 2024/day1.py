from collections import Counter

left = []
right = []

with open("input-day1.txt") as input_file:
    for line in input_file:
        a, b = line.split("   ")
        left.append(int(a))
        right.append(int(b))

left.sort()
right.sort()

result_1 = sum(abs(a-b) for a,b in zip(left, right))
print(result_1)

#1-2
counted = Counter(right)
result_2 = sum([ x*counted[x] for x in left ])
print(result_2)