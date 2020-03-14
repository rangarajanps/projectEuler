function isPrime(number) {

  if ( number < 2 ) return false;
  if ( number == 2 ) return true;
  if ( number % 2 == 0 ) return false;
  for ( var i = 3; i < Math.sqrt(number); i=i+2 ) {
    if ( number % i == 0 ) return false;
  }
  return true;

};

function spiralPrimes() {
  
  var gridNum = 1;
  var sl = 1;
  var i=1;
  var grid = [];
  grid.push(gridNum);

  var num = 0;
  var den = 0;
  
  while ( true ) {
    gridNum = gridNum + 2*i;
    if ( isPrime(gridNum) ) {
      num = num + 1;
    }
    grid.push(gridNum);

    gridNum = gridNum + 2*i;
    if ( isPrime(gridNum) ) {
      num = num + 1;
    }
    grid.push(gridNum);

    gridNum = gridNum + 2*i;
    if ( isPrime(gridNum) ) {
      num = num + 1;
    }
    grid.push(gridNum);

    gridNum = gridNum + 2*i;
    grid.push(gridNum);

    i = i + 1;

    sl = sl+2;
    den = 2*sl - 1;
    //console.log("Side length "+sl+" Grid Number is "+gridNum+" i is "+i);
    //console.log("Numerator is "+num+" Denominator is "+den);
    //console.log("Ratio is "+(num/den));
    if ( (num/den)*100 <= 10 ) {
      //console.log("Generated grid length is "+grid.length);
      //console.log("Generated grid is "+grid);
      console.log("Answer is "+sl);
      return sl;
    }
  }

}

spiralPrimes();
