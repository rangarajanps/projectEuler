def splitTotal(total, coins, minPos):
    if total == 0: return 1
    count = 0
    for i in range(minPos,len(coins)):
        if coins[i] > total: continue
        count += splitTotal(total - coins[i], coins, i)
    return count;

def coinSums(n):
    #Set up the coin values
    coinValues = [1, 2, 5, 10, 20, 50, 100, 200]

    return(splitTotal(n, coinValues, 0))

if __name__=="__main__":
    print(coinSums(50))
