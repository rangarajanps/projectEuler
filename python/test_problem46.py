import unittest
from problem46 import goldbachOtherConjecture

class MyTestCase(unittest.TestCase):
    def test_goldbachOtherConjecture(self):
        self.assertEqual(5777, goldbachOtherConjecture())


if __name__ == '__main__':
    unittest.main()
