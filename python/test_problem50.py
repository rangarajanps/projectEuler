import unittest
from problem50 import consecutivePrimeSum

class MyTestCase(unittest.TestCase):
    def test_consecutivePrimeSum(self):
        testData = {
            100: 41,
            1000: 953,
            1000000: 997651,
        }
        for key in testData:
            self.assertEqual(testData[key], consecutivePrimeSum(key))


if __name__ == '__main__':
    unittest.main()
