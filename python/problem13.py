def sumOfLargeNumbers(tests):
    sum = 0
    for val in tests:
        sum += val
    res = int(str(sum)[:10])
    return res

if __name__=='__main__':
    testcase1 = [37107287533902102798797998220837590246510135740250, 46376937677490009712648124896970078050417018260538]
    print(sumOfLargeNumbers(testcase1))
