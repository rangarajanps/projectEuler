def lexicographicPermutation(remain):
    numbers = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    N = len(numbers)
    permNum = ""

    for i in range(1, N):
        fact = factorial(N - i)
        j = remain // fact
        remain = remain % fact

        permNum = permNum + str(numbers[j])
        numbers.remove(numbers[j])

        if remain == 0:
            break

    for i in range(0, len(numbers)):
        permNum = permNum + str(numbers[i])

    return int(permNum)


def factorial(num):
    if num < 0:
        return 0

    product = 1
    for i in range(2, num + 1):
        product *= i

    return product

if __name__=="__main__":
    print(lexicographicPermutation(699999))
