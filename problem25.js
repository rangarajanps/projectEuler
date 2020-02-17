function countNumberDigits(n) {

  var count = 0;
  while ( n > 0 ) {
    count++;
    n = parseInt(n / 10);
  }
  return count;

}

function digitFibonacci(n) {
  
  var fib0 = 1;
  var fib1 = 2;
  var fib2 = 0;
  var ctr = 3;

  while ( true ) {

    fib2 = fib0 + fib1;
    ctr++;
    if ( countNumberDigits(fib2) >= n ) {
      break;
    }
    //console.log("Last 3 Fibo series numbers are "+fib0+" "+fib1+" "+fib2+" and counter is "+ctr)
    fib0 = fib1;
    fib1 = fib2;
    
  }

  return ctr;
}

//console.log(digitFibonacci(20));
