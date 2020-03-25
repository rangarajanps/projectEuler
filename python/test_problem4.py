import unittest
from problem4 import findMaxPalindrome

class MyTestCase(unittest.TestCase):
    def test_findMaxPalindrome(self):
        testdata = {2:9009, 3:906609}
        for key in testdata:
            self.assertEqual(findMaxPalindrome(key),testdata[key])


if __name__ == '__main__':
    unittest.main()
