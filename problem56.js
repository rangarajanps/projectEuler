function powerDigitSum(base, exponent) {

  var number = [1];
  var carryFwd = 0;
    
  for(var i = 0; i < exponent; i++) {

    var count = number.length;
    
    for(var j = 0; j < count; j++) {

      var digit = number[j] || 0;
      digit = digit*base + carryFwd;    
      if (digit > 9) {
        carryFwd = parseInt(digit/10);
        digit = digit % 10;
      } else {
        carryFwd = 0;
      }
      number[j] = digit;
    }
    if ( carryFwd != 0 ) {
        if ( carryFwd < 10 ) {
          number[number.length] = (number[number.length] || 0) + carryFwd;
        } else {
          number[number.length] = (number[number.length] || 0) + (carryFwd%10);
          number[number.length] = (number[number.length] || 0) + (parseInt(carryFwd/10));
        }
        carryFwd = 0;
    }
    
    //console.log("Number stored as list is "+number);
  }
  //console.log("Number stored as list is "+number.reverse().join(""));

  var sum = 0;
  for(var i = 0; i < number.length; i++) {
    sum += number[i] || 0;
  }

  return sum;
}

function powerfulDigitSum() {
  
  var maxSum = 0;
  var temp = 0;

  for ( var i = 90; i<100; i++) {
    if ( i % 10 == 0 ) continue;
    for ( var j = 90; j<100; j++) {
        temp = powerDigitSum(i, j);
        if ( temp > maxSum ) {
          maxSum = temp;
          //console.log("Storing the max sum for "+i+" power "+j+" value "+maxSum);
        }
    }
    //console.log("Maximum digital sum for "+i+" is "+maxSum);
  }

  console.log("Maximum digital sum is "+maxSum);
  return maxSum;
}

//powerfulDigitSum();
