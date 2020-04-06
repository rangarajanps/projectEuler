import unittest
from problem16 import powerDigitSum

class MyTestCase(unittest.TestCase):
    def test_powerDesignSum(self):
        testData = {
            15: 26,
            128: 166,
            1000: 1366,
        }
        for key in testData:
            self.assertEqual(testData[key], powerDigitSum(key))

if __name__=='__main__':
    unittest.main()
