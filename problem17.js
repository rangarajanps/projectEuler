function convertNumberToWord(number) {

  var oneth = ["","One","TWo","Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven","Twelve", "Thirteen","Fourteen","Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen", "Twenty"];
  var tenth = ["","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninty"];
  var others = ["","","Hundred","","Thousand"];
  
  if ( number <= 19) {
    console.log("Number to word is "+oneth[number]);
    return oneth[number].length;
  }

  var len = 0;
  var original = number;
  var digits = [];
  while (number >=1) {
    digits.push(number%10);
    number = parseInt(number / 10);
    len++;
  }
  digits = digits.reverse();
  console.log("Number of digits in array is "+len+" "+digits);

  var words = 
/*
  var words = [];
  var digit = 0;
  for (var i = 0; i < len; i++) {
    digit = original % Math.pow(10, i+1);
    original = original / 10;
    if ( i == 0) {
      words[i] = oneth[digit];
    } else if ( i == 1) {
      words[i] = tenth[digit];
    }
  }
  console.log("Number in words is "+words);
*/
  var ans = 0;
  for (var i = 0; i < words.length; i++) {
    ans = ans + words[i].length;
  }
  return ans;

}

function numberLetterCounts(limit) {

  var ans = 0;  
  for ( var i = 1; i <= limit; i++) {
    ans = ans + convertNumberToWord(i);
  }
  
  return ans;
}

console.log(numberLetterCounts(21));
