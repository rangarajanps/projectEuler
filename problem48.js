function selfPowers(power, lastDigits) {
  var result = 0;
  var modulo = parseInt("1".concat("0".repeat(lastDigits)));
  
  for (var i = 1; i <= power; i++) {
    var temp = i;
    for (var j = 1; j < i; j++) {
        temp *= i;
        temp %= modulo;
    }
    result += temp;
    result %= modulo;
    
  }
  return result;
}

//console.log(selfPowers(10, 3));
