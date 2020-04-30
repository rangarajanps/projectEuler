import unittest
from problem40 import champernowneConstant

class MyTestCase(unittest.TestCase):
    def test_champernowneConstant(self):
        testData = {
            100: 5,
            1000: 15,
            1000000: 210,
        }
        for key in testData:
            self.assertEqual(testData[key], champernowneConstant(key))


if __name__ == '__main__':
    unittest.main()
