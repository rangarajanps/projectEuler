import unittest
from problem1 import sumOfMultiplesOf3And5

class MyTestCase(unittest.TestCase):
    def test_sumOfMultiplesOf3And5(self):
        testdata = {10:23, 49:543, 1000:233168, 8456:16687353, 19564:89301183}
        for key in testdata:
            self.assertEqual(sumOfMultiplesOf3And5(key),testdata[key])


if __name__ == '__main__':
    unittest.main()
