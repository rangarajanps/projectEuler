import unittest
from problem37 import truncatablePrimeSum

class MyTestCase(unittest.TestCase):
    def test_truncatablePrimeSum(self):
        testData = {
            8: 1986,
            9: 5123,
            10: 8920,
            11: 748317,
        }
        for key in testData:
            self.assertEqual(testData[key], truncatablePrimeSum(key))


if __name__ == '__main__':
    unittest.main()
