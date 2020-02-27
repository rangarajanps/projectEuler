function isBinaryPalindrome(n) {

  var numToTest = n;
  var nBin = [];
  while ( n > 0 ) {
    nBin.push(n%2);
    n = parseInt(n/2);
  }
  //console.log("Binary form of "+numToTest+" is "+nBin);
  
  var max = nBin.length;
  for ( var i = 0; i < max/2; i++) {
    if ( nBin[i] != nBin[max-i-1] ) {
      return false;
    }
  }
  return true;

}

function isPalindrome(num) {

  var nStr = String(num);
  if ( nStr[0] != nStr[nStr.length-1]) return false;

  var numToTest = num;
  var temp = 0;
  while ( numToTest > 9 ) {
    temp = (temp+(numToTest%10))*10;
    numToTest = parseInt(numToTest/10);
  }
  temp = temp + numToTest;
  //console.log("Reversed number for "+num+" is "+temp);
  if ( temp == num ) {
    //console.log("FOUND a palindrome "+num);
    return true;
  }
  return false;
}

function doubleBasePalindromes(n) {

  var pSum = 0;
  var pList = [];
  for ( var i = 1; i <= n; i=i+2 ) {
    if ( isPalindrome(i) && isBinaryPalindrome(i) ) {
      //console.log("FOUND a double base palindrome "+i);
      pList.push(i);
      pSum = pSum + i;
    }
  }
  
  //console.log("Double base Palindrome list is "+pList);
  return pSum;
}

//console.log(doubleBasePalindromes(1000));
