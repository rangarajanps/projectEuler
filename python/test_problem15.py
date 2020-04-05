import unittest
from problem15 import latticePaths

class MyTestCase(unittest.TestCase):
    def test_latticePaths(self):
        testdata = {
            4: 70,
            9: 48620,
            20: 137846528820,
        }
        for key in testdata:
            self.assertEqual(testdata[key], latticePaths(key))


if __name__ == '__main__':
    unittest.main()
