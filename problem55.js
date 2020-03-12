function reverseNumber(number) {
  var rem=0, reversed = 0;
  while(number>0) {
			rem = number%10;
			number = parseInt(number/10);
			reversed = reversed*10+rem;
	}
  return reversed;
}

function isPalindrome(number) {
  
  if (number == reverseNumber(number)) {
    //console.log("Palindrome found "+orignal);
    return true;
  }
  return false;
}

function countLychrelNumbers(limit) {
  
  var temp;
  var lychrelNumList = [];
  var MAX = 50;
  var notLychrel = false;

  for ( var i = 1; i < limit; i++) {

    temp = i;
    notLychrel = false;

    for ( var j = 0; j < MAX; j++) {
      temp = temp+reverseNumber(temp);
      if ( isPalindrome(temp)) {
        notLychrel = true;
        break;
      }
    }
    if ( ! notLychrel ) {
      lychrelNumList.push(i);
    }
  }
  
  return lychrelNumList.length;
}

countLychrelNumbers(1000);
