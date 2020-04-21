import unittest
from problem31 import coinSums

class MyTestcase(unittest.TestCase):
    def test_coinSums(self):
        testData = {
            50:451,
            100:4563,
            150:21873,
            200:73682,
        }
        for key in testData:
            self.assertEqual(testData[key], coinSums(key))

if __name__=="__main__":
    unittest.main()
