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

function isTruncatablePrime(primeList, num) {

  var nStr = String(num);
  var strN = 0;
  if ( nStr.length < 2 ) return false;
  if ( nStr.length > 2 && ( (nStr.indexOf(2) != -1) ||  (nStr.indexOf(5) != -1) ) ) return false;
  for ( var j = 0; j < nStr.length-1; j++ ) {
    strN = parseInt(nStr.substring(j+1,nStr.length));
    //console.log("Original number is "+num+" Checking for truncated number "+strN);
    if ( primeList.indexOf(strN) == -1 ) {
      return false;
    }
    strN = parseInt(nStr.substring(0,nStr.length-j-1));
    //console.log("Original number is "+num+" Checking for truncated number "+strN);
    if ( primeList.indexOf(strN) == -1 ) {
      return false;
    }
  }
  
  return true;
}

function truncatablePrimes(n) {
  var maxN = 800000;

  var primeList = calculatePrimesTo(maxN);
  //console.log("Number of Primes till "+maxN+" to check is "+primeList.length+" and Last is "+primeList[primeList.length-1]);
  var count = 0;
  var sumTrunPrime = 0;

  for ( var i = 0; i < primeList.length; i++ ) {

    if ( isTruncatablePrime(primeList, primeList[i]) ) {
      count++;
      sumTrunPrime = sumTrunPrime + primeList[i];
      //console.log("Found a Truncatable prime "+primeList[i]);
    }
    if ( count == n ) {
      return sumTrunPrime;
    }

  }

  //console.log("Number of Primes found "+count+" and their Sum is "+sumTrunPrime);
  return sumTrunPrime;

}

//truncatablePrimes(11);
