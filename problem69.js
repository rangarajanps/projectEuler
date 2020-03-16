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

function totientMaximum() {
  
  var result = 1;
  var primes = calculatePrimesTo(100);
  //console.log("Primes to check is "+primes);
  var i = 0;
  var limit = 1000000;
 
  while(result* primes[i] < limit){
    result *= primes[i];
    i++;
  }

  //console.log("Max number is "+result+" i is "+i);
  return result;

}

totientMaximum();
