import unittest
from problem29 import lengthDistinctPowerSequence

class MyTestcase(unittest.TestCase):
    def test_lengthDistinctPowerSequence(self):
        testData = {
            15: 177,
            20: 324,
            25: 519,
            30: 755,
        }
        for key in testData:
            self.assertEqual(testData[key], lengthDistinctPowerSequence(key))

if __name__=='__main__':
    unittest.main()
