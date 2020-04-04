import unittest
from problem14 import findLongestCollatzSequence


class MyTestCase(unittest.TestCase):
    def test_findLongestCollatzSequence(self):
        testdata = {14: 9,
                    5847: 3711,
                    46500: 35655,
                    54512: 52527,
                    100000: 77031, }
        for key in testdata:
            self.assertEqual(testdata[key], findLongestCollatzSequence(key))


if __name__ == '__main__':
    unittest.main()
