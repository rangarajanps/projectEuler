function countRightTrianglesWithPeri(n) {

  var count = 0;
  var c = 0;
  for ( var a = 11; a < n/3; a++) {
    for ( var b = a+1; b < n/2; b++) {

      c = Math.sqrt(a*a+b*b);
      if ( ! Number.isInteger(c) ) continue;
      if ( a+b+c == n ) {
        count++;
        //console.log("Found a right triangle for Perimeter "+n+" with sides "+a+" "+b+" "+c);
      }
    }
  }
  return count;
}

function intRightTriangles(n) {
  
  var maxPeri = 0;
  var maxCount = 0;
  var temp = 0;

  for ( var i = 8; i <=n; i++ ) {
    temp = countRightTrianglesWithPeri(i);
    if ( temp >= maxCount ) {//&& maxPeri < i ) {
      maxCount = temp;
      maxPeri = i;
    }
  }
  
  return maxPeri;
}

//console.log(intRightTriangles(120)); // 420
