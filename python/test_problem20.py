import unittest
from problem20 import factorialDigitSum

class MyTestcase(unittest.TestCase):
    def test_factorialDigitSum(self):
        testData = {
            10: 27,
            25: 72,
            50: 216,
            75: 432,
            100: 648,
        }

        for key in testData:
            self.assertEqual(testData[key], factorialDigitSum(key))

if __name__=='__main__':
    unittest.main()
