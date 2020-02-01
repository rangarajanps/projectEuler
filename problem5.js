function findGCD(x, y) {
  while(y) {
    var t = y;
    y = x % y;
    x = t;
  }
  return x;
}

function findLCM(x, y) {
  return ((x * y) / findGCD(x, y));
}

function smallestMult(n) {
  
  var value =  findLCM(2,3);
  for ( var i = 4; i <= n; i++) {
    value = findLCM(value, i);
  }
  
  return value;

}

smallestMult(7);
