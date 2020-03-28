import unittest
from problem7 import generateNthPrimeNumber

class MyTestCase(unittest.TestCase):
    def test_findGenerateNthPrimeNumber(self):
        testdata = {6:13, 10:29, 100:541, 1000: 7919, 10001: 104743}
        for key in testdata:
            self.assertEqual(testdata[key], generateNthPrimeNumber(key) )


if __name__ == '__main__':
    unittest.main()
