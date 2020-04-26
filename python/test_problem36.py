import unittest
from problem36 import doubleBasePalindromeSum

class MyTestcase(unittest.TestCase):
    def test_doubleBasePalindromeSum(self):
        testData = {
            1000: 1772,
            50000: 105795,
            500000: 286602,
            1000000: 872187,
        }
        for key in testData:
            self.assertEqual(testData[key], doubleBasePalindromeSum(key))

if __name__=="__main__":
    unittest.main()
