def fibo(n):
    if n==1:
        return 1
    elif n==2:
        return 1
    else:
        return fibo(n-1)+fibo(n-2)

def sumEvenFibNumbers(limit):
    sum = 0;
    for x in range(1,limit):
        fibValue = fibo(x)
        if fibValue > limit:
            break;
        if fibValue % 2 == 0:
            sum += fibValue
    return sum

if __name__ == '__main__':
    print(sumEvenFibNumbers(60))
