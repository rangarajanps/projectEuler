// uses Sieve of Eratosthenes
function calculatePrimesTo(max) {
    var isPrime = [];
    
    for (var i = 2; i <= max; i++) {
        isPrime[i] = true;
    }

    for (var i = 2; i*i <= max; i++) {
        if (isPrime[i]) {
            for (var j = i; i*j <= max; j++) {
                isPrime[i*j] = false;
            }
        }
    }

    return isPrime;
};

function generatePrimeList(pList) {
    
    var primes = [];    

    for (var i = 2; i <= pList.length; i++) {
        if (pList[i]) {
            primes.push(i);
        }
    }

    return primes;
};

function consecutivePrimeSum(limit) {
  
  var primeList = calculatePrimesTo(limit);
  var primeIntegers = generatePrimeList(primeList);
  var totalPrimes = primeIntegers.length
  
  var count = 0;
  var sum = 0;
  var finalSum = 0;
  var finalCount = 0;

  for (var start = 0; start < totalPrimes; start++) {
    sum = 0;
    count  = 0;
    for (var current = start; current < totalPrimes; current++) {
      var actual = primeIntegers[current];
      sum += actual;
      if ( sum >= limit ) {
        break;
      }
      if ( primeList[sum] ) {
        if ( count > finalCount ) {
          finalCount = count;
          finalSum = sum;
        }
      }
      count++;
    }
  }

  return finalSum;
}

//console.log(consecutivePrimeSum(1000));
