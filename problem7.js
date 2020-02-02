function isPrime(number, series) {

  for (var i = 0; i <= series.length; i++) {
    if ( number % series[i] == 0 ) {
      return false;
    }
  }
  
  return true;
}

function nthPrime(n) {

  var primeSeries = [2,3,5,7,11];
  var ptr = primeSeries.length;

  if ( n <= 5 ) {
    return primeSeries[n];
  }

  var i = 13;
  while ( ptr < n ) {

    if ( isPrime(i, primeSeries) ) {
      primeSeries[ptr] = i;
      ptr++;
    }
    i++;

  }
  
  return primeSeries[primeSeries.length-1];
}

nthPrime(6);
