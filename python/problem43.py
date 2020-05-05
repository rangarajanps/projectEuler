def subDivisiblePandigitNumbers():
    nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    permuts = permutation(nums)
    result = []
    for numSlice in permuts:
        if isSubDivisible(numSlice):
            result.append("".join(str(x) for x in numSlice))
    return result


def permutation(lst):
    # If lst is empty then there are no permutations
    if len(lst) == 0:
        return []

    # If there is only one element in lst then, only
    # one permuatation is possible
    if len(lst) == 1:
        return [lst]

    # Find the permutations for lst if there are
    # more than 1 characters

    l = []  # empty list that will store current permutation

    # Iterate the input(lst) and calculate the permutation
    for i in range(len(lst)):
        m = lst[i]

        # Extract lst[i] or m from the list.  remLst is
        # remaining list
        remLst = lst[:i] + lst[i + 1:]

        # Generating all permutations where m is first
        # element
        for p in permutation(remLst):
            l.append([m] + p)
    return l


def isSubDivisible(numbers):
    # d2d3d4 = 406 is divisible by 2
    # d3d4d5 = 063 is divisible by 3
    # d4d5d6 = 635 is divisible by 5
    # d5d6d7 = 357 is divisible by 7
    # d6d7d8 = 572 is divisible by 11
    # d7d8d9 = 728 is divisible by 13
    # d8d9d10 = 289 is divisible by 17
    d2d3d4 = numbers[1] * 100 + numbers[2] * 10 + numbers[3]
    if d2d3d4 % 2 != 0: return False

    d3d4d5 = numbers[2] * 100 + numbers[3] * 10 + numbers[4]
    if d3d4d5 % 3 != 0: return False

    d4d5d6 = numbers[3] * 100 + numbers[4] * 10 + numbers[5]
    if d4d5d6 % 5 != 0: return False

    d5d6d7 = numbers[4] * 100 + numbers[5] * 10 + numbers[6]
    if d5d6d7 % 7 != 0: return False

    d6d7d8 = numbers[5] * 100 + numbers[6] * 10 + numbers[7]
    if d6d7d8 % 11 != 0: return False

    d7d8d9 = numbers[6] * 100 + numbers[7] * 10 + numbers[8]
    if d7d8d9 % 13 != 0: return False

    d8d9d10 = numbers[7] * 100 + numbers[8] * 10 + numbers[9]
    if d8d9d10 % 17 != 0: return False

    return True


if __name__ == "__main__":
    print(subDivisiblePandigitNumbers())
