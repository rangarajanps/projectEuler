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

function checkIfThreeNumbersAreFormedOfSameDigit(n1, n2, n3) {
  var digits = Array.from(String(n1));
  var nStr2 = Array.from(String(n2));
  var nStr3 = Array.from(String(n3));

  for ( var i = 0; i < digits.length; i++) {
    if ( nStr2.indexOf(digits[i]) == -1 || 
          nStr3.indexOf(digits[i]) == -1 ) {
      return false;
    }
  }

  return true;
}

function primePermutations() {
  var primeList = calculatePrimesTo(9999);
  
  const UNUSUAL_CONST = 3330;
  var p1 = 0;
  var p2 = 0;
  var p3 = 0;
  var result = "";
  
  for ( var i = 0; i < primeList.length; i++) {
    p1 = primeList[i];
    //if ( p1 < 1234 ) continue;
    if ( p1 < 1488 ) continue;
    p2 = p1 + UNUSUAL_CONST;
    p3 = p2 + UNUSUAL_CONST;

    if ( ( primeList.indexOf(p2) != -1 ) && 
          ( primeList.indexOf(p3) != -1 ) ) {

        result = result.concat(p1+""+p2+""+p3)
        if ( checkIfThreeNumbersAreFormedOfSameDigit(p1,p2,p3) ) {
          result = parseInt(result);
          return result;
        }
        result = "";
    }
  }

  return result;
}

console.log(primePermutations());
