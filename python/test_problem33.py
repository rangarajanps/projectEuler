import unittest
from problem33 import lowestDenominatorOfDigitCancelling

class MyTestCase(unittest.TestCase):
    def test_lowestDenominatorOfDigitCancelling(self):
        self.assertEqual(100, lowestDenominatorOfDigitCancelling())


if __name__ == '__main__':
    unittest.main()
