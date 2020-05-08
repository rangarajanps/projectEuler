import math

def selfPowers(uptoN, numOfDigits):
    sum = 1
    for i in range(2,uptoN+1):
        sum += i** i
        #sum += math.pow(i,i)
    value = str(sum)
    #return value[len(value) - numOfDigits - 2:len(value) - 2]
    return value[len(value)-numOfDigits:len(value)]

if __name__=="__main__":
    print(selfPowers(150,5))
