function multiplesOf3and5(number) {
  var sum = 0;
  for (var i=1;i<number; i++) {
    if ( i % 3 == 0 || i % 5 == 0 ) {
      sum = sum + i;
    }
  }
  //console.log("Sum is "+sum);
  return sum;
}
