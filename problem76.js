function countingSummations() {
  var target = 100;
	var digits = [];
  var ways = [];
	for (var k = 0; k < target; k++) {
		digits[k] = k+1;
    ways[k] = 0;
	}
   		
	ways[0] = 1;
  ways[100] = 0;
		 
	for (var i = 0; i < digits.length-1; i++) {
		for (var j = digits[i]; j <= target; j++) {
		  ways[j] += ways[j - digits[i]];
		}
	}

	return ways[ways.length-1];
}

console.log(countingSummations());
