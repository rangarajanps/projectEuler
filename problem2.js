function fiboEvenSum(n) {
  
  var fibSeries = [];
  fibSeries[0] = 1;
  fibSeries[1] = 2;
  var ctr = 2
  for (var i=3; i<=n; i++) {
    fibSeries[ctr] = fibSeries[ctr-1]+fibSeries[ctr-2];
    ctr++;
  }
  //console.log("Fibonacci series is "+fibSeries)

  var sum = 0;
  for (var j=0; j< fibSeries.length; j++) {
    if ( fibSeries[j] % 2 == 0) {
      sum = sum + fibSeries[j];
    }
  }
  //console.log("Sum is "+sum)
  return sum;
}

fiboEvenSum(23);
