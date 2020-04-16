from decimal import *


def find_fraction(i):
    """Finds the fraction of x to the .prec place"""
    getcontext().prec = 1500
    x = Decimal(1) / Decimal(i)
    # print "i = %s x = %s" % (i, x)
    return x


def find_period(x1):
    """Find period for repeating"""
    L = (str(x1))
    length = len(L)
    l_one, l_two, l_three = 0, 0, 0

    if L[2:3] != '0':
        index = L[2:6]
        l_one = 1
    elif L[2:3] == '0' and L[3:4] != '0':
        index = L[3:7]
        l_two = 1
    elif L[2:3] == '0' and L[3:4] == '0':
        index = L[4:8]
        l_three = 1
    else:
        print("There may be an error in your program")

    max_period = 0
    for v in range(3, length):
        if l_one == 1:
            indexi = L[v:(v + 4)]
        elif l_two == 1:
            indexi = L[(v + 1):(v + 5)]
        elif l_three == 1:
            indexi = L[(v + 2):(v + 6)]
        else:
            print("There has been an error")
        if index == indexi:
            max_period = v
            break
    return max_period


def reciprocal_cycles(n):
    place_holder = 0
    for i in range(1, n+1):
        x1 = find_fraction(i)
        y1 = find_period(x1)
        # print "max_period = ", y1
        if y1 > place_holder:
            place_holder = y1
            max_number = i

    return max_number


if __name__ == '__main__':
    print(reciprocal_cycles(10))
