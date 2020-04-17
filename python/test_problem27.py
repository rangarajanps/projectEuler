import unittest
from problem27 import quadraticPrime


class MyTestCase(unittest.TestCase):
    def test_quadraticPrime(self):
        testData = {
            200: -4925,
            500: -18901,
            800: -43835,
            1000: -59231, }
        for key in testData:
            self.assertEqual(testData[key], quadraticPrime(key))


if __name__ == '__main__':
    unittest.main()
