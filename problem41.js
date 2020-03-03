function checkForPandigit(num) {

  var nStr = String(num);
  
  if ( nStr.indexOf(0) != -1 ) return false;
  for ( var i = 1; i <= nStr.length; i++ ) {
    if ( nStr.indexOf(i) == -1 ) return false;
  }
  return true;
}

// uses Sieve of Eratosthenes
function calculatePrimesTo(max) {
    var isPrime = [];
    var primes = [];

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

    for (var i = 2; i <= max; i++) {
        if (isPrime[i]) {
            primes.push(i);
        }
    }

    return primes;
};

function pandigitalPrime(n) {
  
  var minN = parseInt("1".repeat(n));
  var maxN = parseInt("9".repeat(n));

  var result = 0;

  var primeList = calculatePrimesTo(maxN);
  for ( var i = 0; i < primeList.length; i++) {
    if ( primeList[i] < minN ) continue;
    if ( checkForPandigit(primeList[i]) ) {
      if ( result < primeList[i] ) result = primeList[i];
    }
  }

  return result;
}

console.log(pandigitalPrime(4));
