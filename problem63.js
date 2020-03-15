function powerDigitCount(base, exponent) {

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
    
  }
  
  return number.length;
}

function powerfulDigitCounts() {
  
  var exp = 2;
  var base = 2;
  var value = 0;
  var result = 9; //Since there are 9 numbers raised to 1 gives 1 digit
  var reachedMax = false;
  var valueLength = 0;

  while ( exp < 50 ) {

    while ( true ) {
      valueLength = powerDigitCount(base, exp);
      
      if ( valueLength == exp ) {
        result++;
        //console.log("Found a number whose power is same as digit count for "+base+" "+exp+" "+valueLength);
        //exp++;
        base++;
        break;
      }
      if ( valueLength < exp ) {
        base++;
      }
      if ( valueLength > exp ) {
        if ( base < 5 )
          reachedMax = true;
        //console.log("Reached the MAX "+base+" "+exp);
        base = 2;
        exp++;
        break;
      }
    }

    if ( reachedMax ) {
        break;
    }
    
  }
  return result;
}

console.log(powerfulDigitCounts());
