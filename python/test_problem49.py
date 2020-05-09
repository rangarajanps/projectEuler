import unittest
from problem49 import primePermutations

class MyTestCase(unittest.TestCase):
    def test_primePermutations(self):
        self.assertEqual("296962999629", primePermutations())


if __name__ == '__main__':
    unittest.main()
