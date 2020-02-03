function isPrime(number, series) {

  for (var i = 0; series[i] <= Math.sqrt(number); i++) {
    if ( number % series[i] == 0 ) {
      return false;
    }
  }
  
  return true;
}

function primeSummation(n) {

  if ( n==1 ) return 0;
  if ( n==2 ) return 2;
  if ( n==3 ) return 5;
  var primeSeries = [2, 3];
  var ptr = primeSeries.length;
  var value = 5;
  
  for ( var i = 4; i < n; i++) {

    if ( ( i%2 != 0 ) && isPrime(i, primeSeries) ) {
      primeSeries[ptr] = i;
      ptr++;
      value = value + i;
    }    

  }
  
  return value;
}

primeSummation(17);
