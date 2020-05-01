import unittest
from problem41 import largestNLengthPandigitalPrime

class MyTestcase(unittest.TestCase):
    def test_largestNLengthPandigitalPrime(self):
        testData = { 4:4231, 7:7652413}
        for key in testData:
            self.assertEqual(testData[key], largestNLengthPandigitalPrime(key))

if __name__=="__main__":
    unittest.main()
