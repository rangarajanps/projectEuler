function countingSummations(target) {
  
  var coinSizes = [ 2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97 ];
  var	ways = [];
  for (var k = 0; k < target+1; k++) {
    ways[k] = 0;
  }
  ways[0] = 1;
		 
  for (var i = 0; i < coinSizes.length; i++) {
    for (var j = coinSizes[i]; j <= target; j++) {
      ways[j] += ways[j - coinSizes[i]];
    }
  }
	
  return ways[ways.length-1];
}

function primeSummations() {
  //Prime numbers till 100
  var primes = [ 2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97 ];

  var limit = 5000;
  var maxPrimeSum = 0;
  var temp = 0;
  for ( var i = 0; i < primes.length; i++) {
    temp = countingSummations(primes[i]);
    //console.log("Ways in which "+primes[i]+" can be summed is "+temp);
    if ( temp > limit ) {
      maxPrimeSum = primes[i];
      //console.log("Found the max sum prime number");
      break;
    }
  }

  return maxPrimeSum;
}

console.log(primeSummations());
