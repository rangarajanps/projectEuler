function distinctPrimeFactors(targetNumPrimes, targetConsecutive) {

  var consecutiveNumCount = 0;
	var factorCount = 0;
	var startingNumber = 0;
	var numToCheck = 10;
	var foundOne = false;

	while (true) {

	  if (isPrime(numToCheck)) {
		  numToCheck++;
			foundOne = false;
			consecutiveNumCount = 0;
			continue;
		}

		factorCount = getPrimeFactorsCount(numToCheck);
		if (factorCount == targetNumPrimes) {
			if (!foundOne) {
				// console.log("Initializing consecutive count to 1 for " + numToCheck);
				foundOne = true;
				startingNumber = numToCheck;
				consecutiveNumCount = 1;
			} else {
				// console.log("Incrementing consecutive count by 1 for " + numToCheck);
				consecutiveNumCount++;
				if (consecutiveNumCount == targetConsecutive) {
					// console.log("Returning for " + numToCheck);
					return startingNumber;
				}
			}
		} else {
			foundOne = false;
			consecutiveNumCount = 0;
		}
		numToCheck++;
	}
}

function getPrimeFactorsCount(num) {

  var factors = [];
	var factorsCountMap = [];

	for (var i = 2; i <= Math.sqrt(num); i++) {

		if (num % i == 0) {
			var numToTest = num / i;
			if (isPrime(i)) {
         updateFactor(factors, factorsCountMap, i)          
			}
			if (isPrime(numToTest)) {
				updateFactor(factors, factorsCountMap, numToTest)
			}
			if (numToTest == i) {
				continue;
			}

			while (numToTest > 0) {
				if (numToTest % i == 0) {
					if (isPrime(i)) {
						updateFactor(factors, factorsCountMap, i)
					}
					numToTest = numToTest / i;
				} else {
					break;
				}
			}
		}

	}

	return factors.length;
}

function updateFactor(factors, factorsCountMap, numToUpdate) {

  if ( ! factors.includes(numToUpdate) && ! factorsCountMap.includes(numToUpdate)) {
    factors.push(numToUpdate)
    factorsCountMap.push(numToUpdate)
  }
  if ( factors.includes(numToUpdate) && factorsCountMap.includes(numToUpdate)) {
    const index = factors.indexOf(numToUpdate);
    if (index > -1) {
      factors.splice(index, 1);
    }
    factors.push(numToUpdate*numToUpdate)
    factorsCountMap.push(numToUpdate*numToUpdate)
  }
  // if ( ! factors.includes(numToUpdate) && factorsCountMap.includes(numToUpdate)) {
  //   console.log("More than square condition not handled ",numToUpdate)
  // }
}

function isPrime(num) {

	if (num <= 1) {
		return false;
	}
	if (num == 2) {
		return true;
	}

	if (num % 2 == 0) {
		return false;
	}
  
	for (var i = 3; i <= Math.sqrt(num); i += 2) {
		if (num % i == 0) {
			return false;
		}
	}

	return true;
}

console.log(distinctPrimeFactors(2, 2));
