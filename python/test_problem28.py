import unittest
from problem28 import spiralDiagonalsSum

class MyTestCase(unittest.TestCase):
    def test_spiralDiagonalsSum(self):
        testData = {
            101: 692101,
            303: 18591725,
            505: 85986601,
            1001: 669171001,}
        for key in testData:
            self.assertEqual(testData[key], spiralDiagonalsSum(key))


if __name__ == '__main__':
    unittest.main()
