import unittest
from problem24 import lexicographicPermutation

class MyTestcase(unittest.TestCase):
    def test_lexicographicPermutation(self):
        testData = {
            699999: 1938246570,
		    899999: 2536987410,
		    900000: 2537014689,
		    999999: 2783915460,
	    }
        for key in testData:
            self.assertEqual(testData[key], lexicographicPermutation(key))

if __name__=="__main__":
    unittest.main()
