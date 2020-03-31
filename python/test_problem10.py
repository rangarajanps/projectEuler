import unittest
from problem10 import sumOfPrimesTill

class MyTestCase(unittest.TestCase):
    def test_findSumOfPrimesTill(self):
        testdata = {17:41, 2001:277050, 140759:873608362, 2000000:142913828922}
        for key in testdata:
            self.assertEqual(testdata[key], sumOfPrimesTill(key) )


if __name__ == '__main__':
    unittest.main()
