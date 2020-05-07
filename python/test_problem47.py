import unittest
from problem47 import distinctPrimeFactors


class MyTestCase(unittest.TestCase):
    def test_distinctPrimeFactors(self):
        input = [[2, 2], [3, 3], [4, 4]]
        expected = [14, 644, 134043]
        for i in range(0, len(input)):
            self.assertEqual(expected[i], distinctPrimeFactors(input[i][0], input[i][1]))


if __name__ == '__main__':
    unittest.main()
