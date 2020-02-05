function countFactors(number) {

  var counter = 2;
  const isEven = (number%2==0);
  let inc = isEven ? 1 : 2;
  for (var i = 2 ; i <= Math.sqrt(number); i=i+inc) {
    if ( number % i == 0 ) {
      counter=counter+2;
    }
  }

  return counter;
}


function divisibleTriangleNumber(n) {
  
  var divCount = 0;
  var triangNum = 1;
  var itr = 2;
  while ( true ) {
    triangNum = triangNum + itr;
    itr++;
    if ( triangNum <= n*n ) {
      continue;
    } else {
      break;
    }
  }
  
  for (var i = itr; divCount <=n; i++) {
    divCount = countFactors(triangNum);
    
    //console.log("Triangle Number, DivCount and MaxDiv is "+triangNum+" "+divCount);
    
    if ( divCount > n) {
      return triangNum;
    }
    triangNum = triangNum + i;
  }
  
}

divisibleTriangleNumber(5);
