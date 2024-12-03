#!/home/amitav/anaconda3/bin/python

f = open("./2.in").read().strip()
lines = f.split("\n")

ans = 0
for line in lines:
    xs1 = list(map(int, line.split()))
    good = False
    for k in range(len(xs1)):
        xs = xs1[:k] + xs1[k+1:]
        inc_or_dec = (xs == sorted(xs)) or (xs == sorted(xs, reverse=True))
        ok = True
        for i in range(len(xs)-1):
            diff = abs(xs[i] - xs[i+1])
            if not 1 <= diff <= 3:
                ok = False
        if inc_or_dec and ok:
            good = True
    if good:
        ans += 1

print(ans)
