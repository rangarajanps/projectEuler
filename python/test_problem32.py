import unittest
from problem32 import pandigitNumbers

class MyTestCase(unittest.TestCase):
    def test_pandigitNumbers(self):
        self.assertEqual(45228, pandigitNumbers())


if __name__ == '__main__':
    unittest.main()
