def sumOfMultiplesOf3And5(limit):
    sum = 0
    for x in list([n for n in range(1,limit) if n % 3 == 0 or n % 5 == 0]):  #numbers:
        sum += x
    return(sum)

if __name__ == '__main__':
    print(sumOfMultiplesOf3And5(10))
