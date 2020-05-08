import unittest
from problem48 import selfPowers


class MyTestCase(unittest.TestCase):
    def test_selfPowers(self):
        input = [[10, 3], [150, 5], [673, 7], [1000, 10]]
        expected = ["317", "29045", "2473989","9110846700"]
        for i in range(0, len(input)):
            self.assertEqual(expected[i], selfPowers(input[i][0], input[i][1]))


if __name__ == '__main__':
    unittest.main()
