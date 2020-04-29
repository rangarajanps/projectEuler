import unittest
from problem39 import intRightTriangles

class MyTestCase(unittest.TestCase):
    def test_intRightTriangles(self):
        testData = {
            500: 420,
            800: 720,
            900: 840,
            1000: 840,
        }
        for key in testData:
            self.assertEqual(testData[key], intRightTriangles(key))


if __name__ == '__main__':
    unittest.main()
