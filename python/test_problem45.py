import unittest
from problem45 import triPentaHexa

class MyTestCase(unittest.TestCase):
    def test_triPentaHexa(self):
        self.assertEqual(1533776805, triPentaHexa())


if __name__ == '__main__':
    unittest.main()