import unittest
from problem43 import subDivisiblePandigitNumbers

class MyTestCase(unittest.TestCase):
    def test_subDivisiblePandigitNumbers(self):
        expected = ['1406357289', '1430952867', '1460357289', '4106357289', '4130952867', '4160357289'];
        self.assertEqual(expected, subDivisiblePandigitNumbers())


if __name__ == '__main__':
    unittest.main()
