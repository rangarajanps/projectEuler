function isPalindrome(number) {

  var orignal = number;
  var rem=0, reversed = 0;
  while(number>0) {
	  rem = number%10;
		number = parseInt(number/10);
		reversed = reversed*10+rem;
	}

  if (orignal == reversed) {
    //console.log("Palindrome found "+orignal);
    return true;
  }
  return false;
}

function largestPalindromeProduct(n) {
  
  var num1 = parseInt("9".repeat(n));
  var num2 = num1;
  var maxnum = 9;
  for (var x = 1; x < n; x++) {
    maxnum = maxnum*10
  }

  var value = 0;
  for (var i = num1; i > maxnum; i--) {
    for (var j = num2; j > maxnum; j--) {

      value = i * j;
      //console.log("Iteration with "+i+" * "+j+" Value is "+value);
      if (isPalindrome(value)) {
        //console.log("Value is "+value);
        return value;
      } 
    }
  } 

}

largestPalindromeProduct(3);
