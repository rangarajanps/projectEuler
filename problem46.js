// uses Sieve of Eratosthenes
function calculatePrimesTo(max) {
    var isPrime = [];
    var primes = [];

    for (var i = 2; i <= max; i++) {
        isPrime[i] = true;
    }

    for (var i = 2; i*i <= max; i++) {
        if (isPrime[i]) {
            for (var j = i; i*j <= max; j++) {
                isPrime[i*j] = false;
            }
        }
    }

    for (var i = 2; i <= max; i++) {
        if (isPrime[i]) {
            primes.push(i);
        }
    }

    return primes;
};

function checkGoldbach(num, nList) {
  var temp = 0;
  for ( var i = 0; i<nList.length; i++) {
    temp = num - nList[i];
    if ( temp % 2 != 0 ) continue;
    if ( Number.isInteger(Math.sqrt(temp/2)) ) {
      //console.log(num+" can be written as twice the square");
      return true;
    }
  }
  return false;

}


function goldbachsOtherConjecture() {

  var limit = 6000;
  var primeList = calculatePrimesTo(limit);
  //console.log("Numbers to check is "+(parseInt((limit)/2) - primeList.length))
  
  for ( var i = 9; i < limit; i++ ) {

    if ( i % 2 == 0 || primeList.indexOf(i) != -1 ) {
      continue; // Even composite or prime, skip
    }
    
    const primeSubList = primeList.filter(it => it < i);
    //console.log("Number of primes to verify "+primeSubList);
    if ( ! checkGoldbach(i, primeSubList) ) {
      console.log("Found a number which breaks the Goldbach's other conjecture "+i)
      return i;
    }
    
   }

}

//goldbachsOtherConjecture();
