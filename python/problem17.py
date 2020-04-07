def numberLetterCount(limit):
    sum = 0
    for i in range(1,limit+1):
        word = convertNumberToLetter(i)
        word = word.replace(" ", "")
        sum += len(word)
    return sum

def convertNumberToLetter(n):
    ones = ["", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"]
    tens = ["", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"]

    if n < 20 :
        return ones[n]
    elif n < 100:
        tenth = n // 10
        return tens[tenth] + " " + ones[n % 10]
    elif n < 1000:
        hundreth = n // 100
        if n % 100 == 0:
            return ones[hundreth] + " hundred "
        return ones[hundreth] + " hundred and " + convertNumberToLetter(n % 100)
    elif n < 10000:
        thousanth = n // 1000
        if n % 1000 == 0:
            return ones[thousanth] + "thousand "
        return ones[thousanth] + "thousand and " + convertNumberToLetter(n % 1000)
    else:
        million = n // 1000000
        if n % 1000000 == 0:
            return ones[million] + "million "
        return ones[million] + "thousand and " + convertNumberToLetter(n % 1000000)

if __name__=='__main__':
    print(numberLetterCount(150))
