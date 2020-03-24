import unittest
from problem2 import sumEvenFibNumbers

class MyTestCase(unittest.TestCase):
    def test_sumEvenFibNumbers(self):
        testdata = {10: 10, 60: 44, 1000: 798, 100000: 60696, 4000000: 4613732}
        for key in testdata:
            self.assertEqual(sumEvenFibNumbers(key),testdata[key])


if __name__ == '__main__':
    unittest.main()
