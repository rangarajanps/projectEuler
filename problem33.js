function digitCancellingFractions() {
  
  var num = 0;
  var den = 0;
  var numProd = 1;
  var denProd = 1;  

  for (var i=10; i <= 99; i++ ) {
    for (var j=i+1; j <= 99; j++ ) {
      if ( i%10 == parseInt(j/10) ) {
        num = parseInt(i/10);
        den = j%10;
        if ( num / den == (i/j) ) {
          numProd = numProd * num;
          denProd = denProd * den;
        }
      }
    }
  }

  num = numProd/denProd;
  den = 1;
  while ( num < 1 ) {
    num = num*10;
    den = den*10;
  }
  
  return den;
}

digitCancellingFractions();
