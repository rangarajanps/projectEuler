import unittest
from problem18_67 import maximumPathSum

class MyTestCase(unittest.TestCase):
    def test_maximumPathSum(self):

        m1 = [[3, 0, 0, 0],
		     [7, 4, 0, 0],
		     [2, 4, 6, 0],
		     [8, 5, 9, 3]]
        m2 = create_matrix_from_file("matrix2.txt")
        m3 = create_matrix_from_file("matrix3.txt")

        self.assertEqual(23, maximumPathSum(m1))
        self.assertEqual(1074, maximumPathSum(m2))
        self.assertEqual(7273, maximumPathSum(m3))

def create_matrix_from_file(filename):
    file1 = open(filename, 'r')
    lines = file1.readlines()
    file1.close()

    a = [[0 for x in range(len(lines))] for y in range(len(lines))]

    for i in range(0, len(lines)):
        numbers = lines[i].split(" ")
        for j in range(0,len(numbers)):
            val= int(numbers[j])
            a[i][j] = val
    return a


if __name__ == '__main__':
    unittest.main()
