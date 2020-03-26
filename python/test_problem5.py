import unittest
from problem5 import findSmallestFactor

class MyTestCase(unittest.TestCase):
    def test_findSmallestFactor(self):
        testdata = {5:60, 7:420, 10:2520, 13:360360, 20:232792560}
        for key in testdata:
            self.assertEqual(findSmallestFactor(key),testdata[key])


if __name__ == '__main__':
    unittest.main()

