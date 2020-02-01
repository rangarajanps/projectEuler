function isPrime(number) {
  
  for (var i=2; i<= number/2; i++) {
    if ( number % i == 0 ) {
      return false;
    }
  }

  return true;
}

function largestPrimeFactor(number) {

  if ( isPrime(number)) return number;

  var primeFactors = [];
  var ptr = 0;
  for (var i=2; i<= number/2; i++) {
    if ( number % i == 0 && isPrime(i)) {
      primeFactors[ptr] = i;
      ptr++;
    }
  }

  //console.log("Prime factor series is "+primeFactors)
  //console.log("Largest Prime factor is "+primeFactors[primeFactors.length-1])
  return primeFactors[primeFactors.length-1];
}

largestPrimeFactor(13195);
