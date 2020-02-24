function splitTotal(total, coins, minPos) {
    if (total === 0) return 1;
    var count = 0;
    for (var i = minPos; i < coins.length; i++) {
        if (coins[i] > total) continue;
        count += splitTotal(total - coins[i], coins, i);
    }
    return count;
};

function coinSums(n) {

  //Set up the coin values
  var coinValues = [1, 2, 5, 10, 20, 50, 100, 200];

  return(splitTotal(n, coinValues, 0));
}

console.log(coinSums(10));
