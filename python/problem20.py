def factorialDigitSum(n):
	return findSum(factorial(n))

def factorial(n):
	fact = 2
	for i in range(3,n+1):
		fact = fact * i
	return fact

def findSum(n):
	sum = 0
	while n > 0:
		sum = sum + n%10
		n = n // 10
	return sum


if __name__=="__main__":
    print(factorialDigitSum(10))
