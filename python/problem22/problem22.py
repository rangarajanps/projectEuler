def nameScores(wordList):
    sum=0
    wordList.sort()
    for i in range(0, len(wordList)):
        sum += (i+1)*calculateScore(wordList[i])
    return sum

def calculateScore(inp):
    score = 0
    for i in range(0, len(inp)):
        score += ord(inp[i])-64
    return score

if __name__=='__main__':
    test1 = ['I', 'REPEAT', 'THIS', 'IS', 'ONLY', 'A', 'TEST'] #['THIS', 'IS', 'ONLY', 'A', 'TEST']
    print(nameScores(test1))
