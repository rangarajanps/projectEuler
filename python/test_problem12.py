import unittest
from problem12 import findHighlyDivisibleTriangularNumber

class MyTestCase(unittest.TestCase):
    def test_findHighlyDivisibleTriangularNumber(self):
        testData = { 5:28,
                23:  630,
		            167: 1385280,
		            374: 17907120,
		            500: 76576500}
        for key in testData:
            self.assertEqual(testData[key], findHighlyDivisibleTriangularNumber(key))

if __name__ == '__main__':
    unittest.main()
