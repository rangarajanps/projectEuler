import unittest
from problem30 import digitNPower

class MyTestcase(unittest.TestCase):
    def test_digitNPower(self):
        testData = {2:0,3:1301, 4:19316,5:443839}
        for key in testData:
            self.assertEqual(testData[key],digitNPower(key),"Test for Euler Problem30")

if __name__=="__main__":
    unittest.main()
