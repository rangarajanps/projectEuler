import unittest
from problem25 import nDigitFibonacciNumber

class MyTestCase(unittest.TestCase):
    def test_nDigitFibonacciNumber(self):
        testData = {5:21, 10:45, 15:69, 20:93, }
        for key in testData:
            self.assertEqual(testData[key], nDigitFibonacciNumber(key))


if __name__ == '__main__':
    unittest.main()
