import unittest
from problem23 import nonAbundantSum

class MyTestcase(unittest.TestCase):
    def test_nonAbundantSum(self):
        testData = {10000:3731004, 15000:4039939, 20000:4159710, 28123:4179871 }
        for key in testData:
            self.assertEqual(testData[key],nonAbundantSum(key))

if __name__=='__main__':
    unittest.main()
