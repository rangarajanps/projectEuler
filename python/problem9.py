import math

def findSpecialPythagoreanTriplet(limit):
    for a in range(1,(limit//2)+1):
        for b in range(a+1,limit//2):
            c = math.sqrt(a*a+b*b)
            if c.is_integer():
                c = int(c)
                //print(a, b, c, (a + b + c))
                if a+b+c == limit: return a*b*c

if __name__=='__main__':
    print(findSpecialPythagoreanTriplet(24))
