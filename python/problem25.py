def nDigitFibonacciNumber(n):
    if n ==1: return 1
    fiboSeries = {0:1, 1:1} 
    i=2
    while True:
        fiboSeries[i] = fiboSeries.get(i-1)+fiboSeries.get(i-2)
        if countDigits(fiboSeries.get(i)) == n:
            i+=1
            return i
        i+=1

def countDigits(num):
    count=0
    while num > 0:
        count+=1
        num = num // 10
    return count

if __name__=="__main__":
    print(nDigitFibonacciNumber(5))
