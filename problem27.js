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

function isPrime(n, primes) {
    var maxCheck = Math.sqrt(Math.abs(n));
    for (var i = 0; primes[i] <= maxCheck; i++) {
        if (n % primes[i] === 0) {
            return false;
        }
    }
    return true;
};

var quadratic = function(a, b, n) {
    return n*n + a*n + b;
};

// As soon as n == b the quadratic is not prime (b is present in every term).
// Given this, the largest value possible from the quadratic is when a and b are at max and n == b - 1
function quadraticPrimes(limit) {

  var primes = calculatePrimesTo(Math.sqrt(quadratic(limit, limit, limit-1)));
    var maxN = 0;
    var maxA, maxB;

    for (var a = -limit; a <= limit; a++) {
        for (var b = -limit; b <= limit; b++) {
            for (var n = 0; n < b && isPrime(Math.abs(quadratic(a, b, n)), primes); n++) {} // empty for body
            if (n > maxN) {
                maxN = n;
                maxA = a;
                maxB = b;
            }
        }
    }

    return maxA * maxB;

}

quadraticPrimes(200);
