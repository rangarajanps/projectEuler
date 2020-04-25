import unittest
from problem35 import circularPrimeCount

class MyTestCase(unittest.TestCase):
    def test_circularPrimeCount(self):
        testData = {
            100: 13,
            100000: 43,
            250000: 45,
            500000: 49,
            750000: 49,
            1000000: 55
        }
        for key in testData:
            self.assertEqual(testData[key], circularPrimeCount(key), "Test for Euler Problem35")


if __name__ == '__main__':
    unittest.main()
