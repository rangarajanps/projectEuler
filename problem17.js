function convertNumberToWord(number) {

  var oneth = ["","One","Two","Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven","Twelve", "Thirteen","Fourteen","Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen", "Twenty"];
  var tenth = ["","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"];
  var others = ["","","Hundred","","Thousand"];
  
  if ( number <= 19) {
    //console.log("Number to word is "+oneth[number]);
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
  //digits = digits.reverse();
  //console.log("Number of digits in array is "+len+" "+digits);

  var words;
  if ( original > 19 && original < 100 ) {
    words = tenth[digits[1]]+" "+oneth[digits[0]];
  }

  var digit = 0, remind = 0;
  if (original == 100) {
    words = "One hundred";
  }
  if ( original > 100 && original < 1000 ) {
    digit = parseInt(original / 100);
    remind = original % 100;
    if ( remind > 19) {
      words = oneth[digit] +" hundred and "+tenth[digits[1]]+" "+oneth[digits[0]];
    } else {
      words = oneth[digit] +" hundred and "+oneth[remind];
    }
    
  }

  if (original == 1000) {
    words = "One Thousand";
  }
  if ( original > 1000 && original < 10000 ) {
    digit = parseInt(original / 1000);
    remind = original % 1000;
    if ( remind > 19) {
      words = oneth[digit] +" Thousand "+oneth[digits[2]] +" hundred and "+tenth[digits[1]]+" "+oneth[digits[0]];
    } else {
      words = oneth[digit] +" Thousand "+oneth[digits[2]] +" hundred and "+oneth[remind];
    }
    console.log("Number in words for "+original+" is "+words);
  }
  //console.log("Number in words for "+original+" is "+words);
  
/*
  
  var digit = 0;
  for (var i = 0; i < len; i++) {
    digit = original % Math.pow(10, i+1);
    //original = original / 10;
    if ( i == 0) {
      words[i] = oneth[digit];
    } else if ( i == 1) {
      words[i] = tenth[digit];
    }
  }
  
*/

  var ans = 0;
  for (var i = 0; i < words.length; i++) {
    words = words.replace(/\s/g,"");
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

console.log(numberLetterCounts(1000));
