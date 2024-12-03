#!/home/amitav/anaconda3/bin/python

L1 = []
L2 = []

f = open("./1.in")
for line in f:
    a, b = [int(x) for x in line.split("   ")]
    L1.append(a)
    L2.append(b)
f.close()

list1 = L1.copy()
list2 = L2.copy()

assert len(list1) == len(list2)

d = 0
while not len(list1) == 0:
    m1, m2 = min(list1), min(list2)
    d += abs(m1 - m2)
    list1.remove(m1)
    list2.remove(m2)

print("Part 1:", d)

list1 = L1.copy()
list2 = L2.copy()

s = 0
for n in list1:
    c = list2.count(n)
    s += n * c

print("Part 2:", s)
