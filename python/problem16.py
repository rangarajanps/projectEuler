import math

def powerDigitSum(exp):
    val = str(int(math.pow(2, exp)))
    sum = 0
    for i in val:
        sum += int(i)

    return sum

if __name__=='__main__':
    print(powerDigitSum(15))
