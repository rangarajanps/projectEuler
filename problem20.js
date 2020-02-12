function factorial(n) {
  var fact = 1;
  for (var i = 1; i<=n; i++) {
    fact = fact * i;
  }
  return fact;
}

function findSum(n) {

  var sum = 0;
  while ( n > 0 ) {
    sum = sum + n%10;
    n = parseInt(n / 10);
  }
  return sum;

}

function convertNumToArray(n) {

  var num = [];
  while ( n > 0 ) {
    num.push(n%10);
    n = parseInt(n / 10);
  }
  return num;

}

function sumFactorialDigits(n) {

  if ( n <= 10 ) {
    return findSum(factorial(n));
  }

  var fact = convertNumToArray(factorial(10));
  //console.log("Factorial of 10 is "+fact);

  var temp = 0;
  var carryFwd = 0;
  for (var i = 11; i <= n; i++) {
    for ( var j = 0; j < fact.length; j++) {
      temp = carryFwd + fact[j]*i;
      fact[j] = temp % 10;
      carryFwd = parseInt(temp / 10 );
    }
    
    var carryFwdInArr = [];
    if ( carryFwd > 0 ) {
      carryFwdInArr = convertNumToArray(carryFwd);
      for ( var x = 0; x < carryFwdInArr.length; x++ ) {
        fact[j] = carryFwdInArr[x]; 
        j++;
      }
      carryFwd = 0;
    }
    //console.log("Factorial of "+i+" is "+fact);
  }

  //console.log("Factorial of "+i+" is "+fact);
  var ans = 0;
  for (var i=0; i < fact.length; i++) {
    ans = ans + fact[i];
  }
  return ans;
  
}

sumFactorialDigits(25);
