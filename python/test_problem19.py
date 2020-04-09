import unittest
from problem19 import countingSundays

class MyTestcase(unittest.TestCase):
    def test_countingSundays(self):
        fromYears = [1943,1995,1901]
        toYears = [1946,2000,2000]
        numOfSundays =[6,10,171]

        for i in range(0,len(fromYears)):
            actual = countingSundays(fromYears[i], toYears[i])
            self.assertEqual(actual,numOfSundays[i])

if __name__=="__main__":
    unittest.main()
