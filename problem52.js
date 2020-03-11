function checkIfTwoNumbersContainSameDigits(num1, num2) {
  var strN1 = String(num1);//.split('');
  var strN2 = String(num2);
  for ( var i = 0; i < strN1.length; i++) {
    if ( strN2.indexOf(strN1[i]) == -1 ) {
      return false;
    }
  }
  return true;
}

function permutedMultiples() {
  
  var n2 = 0;
  var result = 0;
  var doesNotContainSameDigit = true;
  for ( var i = 123456; ; i++ ) {
    doesNotContainSameDigit = false;
    for ( var j = 2; j < 7; j++) {
      n2 = i*j;
      if ( ! checkIfTwoNumbersContainSameDigits(i, n2) ) {
        doesNotContainSameDigit = true;
        break;
      }
    }
    if ( ! doesNotContainSameDigit ) {
      //console.log("Found a Permuted multiples "+i);
      result = i;
      break;
    }
  }
  
  return result;
}


permutedMultiples();
