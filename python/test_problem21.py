import unittest
from problem21 import findAmicableNumberSum

class MyTestcase(unittest.TestCase):
    def test_findAmicableNumberSum(self):
        testData = {
            1000: 504,
            2000: 2898,
            5000: 8442,
            10000: 31626,
        }

        for key in testData:
            self.assertEqual(testData[key], findAmicableNumberSum(key))

if __name__=='__main__':
    unittest.main()
