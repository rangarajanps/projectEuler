function isTriangular(n) {
  let p = (Math.sqrt(8*n+1) -1 ) / 2
  return p === Math.floor(p)
}

function isPentagonal(n) {
  let p = (1 + Math.sqrt(1 + 24*n)) / 6
  return p === Math.floor(p)
}

function isHexagonal(n) {
  let p = (1 + Math.sqrt(1 + 8*n)) / 4
  return p === Math.floor(p)
}

function triPentaHexa(n) {

  var numToTest = 285;
  var result = 0;
  while ( true ) {

    numToTest++;
    result = numToTest * (2 * numToTest - 1);
 
    if (isPentagonal(result) && isHexagonal(result) ) {
        //console.log(result+" is Triangular, pentagonal and hexagonal ");
        return result;
    }
    
  }

}

triPentaHexa(40756);
