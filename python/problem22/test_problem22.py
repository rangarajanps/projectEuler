import unittest
from problem22 import nameScores

class MyTestcase(unittest.TestCase):
    def test_nameScores(self):
        test1 = ['THIS', 'IS', 'ONLY', 'A', 'TEST']
        test2 = ['I', 'REPEAT', 'THIS', 'IS', 'ONLY', 'A', 'TEST']
        test3 = create_list_from_file('p022_names.txt')
        self.assertEqual(791, nameScores(test1))
        self.assertEqual(1468, nameScores(test2))
        self.assertEqual(871198282, nameScores(test3))

def create_list_from_file(filename):
    file1 = open(filename, 'r')
    lines = file1.readlines()
    file1.close()

    wordList = []

    for i in range(0, len(lines)):
        words = lines[i].split(",")
        for j in range(0,len(words)):
            wordList.append(words[j])

    return wordList
