import unittest
from problem6 import findSumSquareDifference

class MyTestCase(unittest.TestCase):
    def test_findSmallestFactor(self):
        testdata = {10:2640, 20:41230, 100:25164150}
        for key in testdata:
            self.assertEqual(findSumSquareDifference(key),testdata[key])


if __name__ == '__main__':
    unittest.main()

