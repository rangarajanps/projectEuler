import unittest
from problem9 import findSpecialPythagoreanTriplet

class MyTestCase(unittest.TestCase):
    def test_findSpecialPythagoreanTriplet(self):
        testdata = {24:480, 120:49920, 1000:31875000}
        for key in testdata:
            self.assertEqual(testdata[key], findSpecialPythagoreanTriplet(key) )


if __name__ == '__main__':
    unittest.main()
