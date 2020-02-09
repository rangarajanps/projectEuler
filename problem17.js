function convert(n) {

  var units = [ "", "One", "Two", "Three", "Four",
"Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
"Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen",
"Eighteen", "Nineteen" ];

var tens = ["", "",	"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", 	"Ninety"];

	if (n < 20) {
		return units[n];
	}

	if (n < 100) {
		return tens[parseInt(n / 10)]+ " " + units[n % 10];
	}

	if (n < 1000 ) {

    if ( n % 100 == 0 ) {
      return units[parseInt(n / 100)] + " Hundred";
    } else {
		  return units[parseInt(n / 100)] + " Hundred" + " and " + convert(n % 100);
    }
	}

	if (n < 100000) {
		return convert(n / 1000) + " Thousand" + " " +  convert(n % 1000);
	}

	if (n < 10000000) {
		return convert(n / 100000) + " Lakh" + " "  + convert(n % 100000);
	}

	return convert(n / 10000000) + " Crore" + " "  + convert(n % 10000000);
}

function convertNumberToWord(number) {

  var words = convert(number);
  //console.log("Number in words for "+number+" is "+words);
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

numberLetterCounts(150);
