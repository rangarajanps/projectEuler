import unittest
from problem8 import findLargestProduct

class MyTestCase(unittest.TestCase):
    def test_findfindLargestProduct(self):
        testdata = {4:5832, 13:23514624000}
        for key in testdata:
            self.assertEqual(testdata[key], findLargestProduct(key) )


if __name__ == '__main__':
    unittest.main()
