function powerDigitSum(exponent) {

  var number = [1];

  for(var i = 0; i < exponent; i++) {

    var carryFwd = 0;
    var count = number.length+1;
    
    for(var j = 0; j < count; j++) {

      var digit = number[j] || 0;
      digit = digit*2 + carryFwd;      
      if (digit > 9) {
        digit = digit - 10;
        carryFwd = 1;
      } else {
        carryFwd = 0;
      }
      number[j] = digit;
    }
    
  }

  var sum = 0;
  for(var i = 0; i < exponent; i++) {
    sum += number[i];
  }

  return(sum);
}

console.log(powerDigitSum(15));
