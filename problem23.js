function isAbundant(number) {

  if ( number < 12 ) {
    return false;
  }    

  var sum = 0;
  var possibleFactor = 1;
    var sqrt = parseInt(Math.sqrt(number));
    while (possibleFactor <= sqrt) {
        if (number % possibleFactor == 0) {
            sum = sum + possibleFactor;

            var otherPossibleFactor = number / possibleFactor;
            if (otherPossibleFactor > possibleFactor && otherPossibleFactor != number ) {
                sum = sum + otherPossibleFactor;
            }
        }
        if ( sum > number ) {
          return true;
        }
        possibleFactor++;
    }

  if ( sum > number ) {
    return true;
  }

  return false;

}

function generateNAbundantNumber(n) {

  var abunNums = [12,18,20,24,30,36,40,42,48];
  for (var i = 50; i <=n; i++ ) {
    if ( isAbundant(i) ) {
      abunNums.push(i);
    }
  }
  return abunNums;

}

function sumOfNonAbundantNumbers(n) {
  
  var abunList = generateNAbundantNumber(n);
  
  var sum = 0;
  
  var isSumOfAbundant = false;
  var abundantNumber1 ;
  
  for ( var testNum = 1; testNum <= n; testNum++ ) {

    isSumOfAbundant = false;

    var j = abunList.length - 1;
    for ( var i = 0 ; abunList[i] < testNum; i++ ) {

      abundantNumber1 = abunList[i];  
      if ( abundantNumber1 > testNum ) {
        break;
      }

      var abundantNumber2 = abunList[j];
      while (( j > 0 ) && ( abundantNumber1 + abundantNumber2 > testNum )) {          
          abundantNumber2 = abunList[--j];
      }
        
      if ( testNum == abundantNumber1 + abundantNumber2 ) {
        //console.log("Identified one "+testNum+" "+abunList[i] +" "+ abunList[j]);
        isSumOfAbundant = true;
        break;
      }
    }
    
    if ( ! isSumOfAbundant ) {
      sum = sum + testNum;
      //console.log("Sum is "+sum+" for i "+i+" testNum "+testNum);
    }
    
  }

  return sum;

}

//sumOfNonAbundantNumbers(28123);
//console.log("Sum of Non Abundant numbers are "+sumOfNonAbundantNumbers(10000));
