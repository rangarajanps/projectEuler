import unittest
from problem34 import digitFactorialsSum

class MyTestCase(unittest.TestCase):
    def test_digitFactorialsSum(self):
        self.assertEqual(40730, digitFactorialsSum())


if __name__ == '__main__':
    unittest.main()
