import math


def lengthDistinctPowerSequence(limit):
    powerSet = set()
    for i in range(2, limit + 1):
        for j in range(2, limit + 1):
            powerSet.add(math.pow(i, j))

    return len(powerSet)


if __name__ == '__main__':
    print(lengthDistinctPowerSequence(5))
