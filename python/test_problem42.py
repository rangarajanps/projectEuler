import unittest
from problem42 import countCodedTriangleNumbers

class MyTestCase(unittest.TestCase):
    def test_countCodedTriangleNumbers(self):
        testData = {
            1400: 129,
            1500: 137,
            1600: 141,
            1786: 162,
        }
        for key in testData:
            self.assertEqual(testData[key], countCodedTriangleNumbers(key))


if __name__ == '__main__':
    unittest.main()
