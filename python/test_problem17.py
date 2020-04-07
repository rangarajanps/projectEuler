import unittest
from problem17 import numberLetterCount


class MyTestCase(unittest.TestCase):
    def test_numberLetterCount(self):
        testData = {
            5: 19,
            150: 1903,
            1000: 21124
        }
        for key in testData:
            self.assertEqual(testData[key], numberLetterCount(key))

if __name__ == '__main__':
    unittest.main()
