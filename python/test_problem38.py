import unittest
from problem38 import largestPandigitalMultiples

class MyTestcase(unittest.TestCase):
    def test_largestPandigitalMultiples(self):
        self.assertEqual(932718654, largestPandigitalMultiples())


if __name__ == '__main__':
    unittest.main()
