import unittest
from problem3 import largestPrimeFactor

class MyTestCase(unittest.TestCase):
    def test_largestPrimeFactor(self):
        testdata = {2:2, 3:3, 5:5, 7: 7, 13195: 29, 600851475143: 6857}
        for key in testdata:
            self.assertEqual(largestPrimeFactor(key),testdata[key])


if __name__ == '__main__':
    unittest.main()
