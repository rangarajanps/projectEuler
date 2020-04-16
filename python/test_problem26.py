import unittest
from problem26 import reciprocal_cycles

class MyTestCase(unittest.TestCase):
    def test_reciprocal_cycles(self):
        testData = {
            700:  659,
		    800:  743,
		    900:  887,
		    1000: 983,}
        for key in testData:
            self.assertEqual(testData[key], reciprocal_cycles(key))


if __name__ == '__main__':
    unittest.main()
