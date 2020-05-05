import unittest
from problem44 import pentagonNumbers

class MyTestCase(unittest.TestCase):
    def test_pentagonNumbers(self):
        self.assertEqual(5482660.0, pentagonNumbers())


if __name__ == '__main__':
    unittest.main()