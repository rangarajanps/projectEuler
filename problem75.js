function gcd(x, y) {
  while(y) {
    var t = y;
    y = x % y;
    x = t;
  }
  return x;
}

function singularIntRightTriangles() {

  var limit = 1500000;
  var triangles = [];
  for ( var i = 0; i < limit; i++ ) {
     triangles[i] = 0;
  }
 
  var result = 0;
  var mlimit = parseInt(Math.sqrt(limit / 2));
  
  for (var m = 2; m < mlimit; m++) {

    for (var n = 1; n < m; n++) {
        
        if (((n + m) % 2) == 1 && gcd(n, m) == 1) {
          
            var a = m * m + n * n;
            var b = m * m - n * n;
            var c = 2 * m * n;
            var p = a + b + c;
            while(p <= limit){
                triangles[p]++;
                //console.log(a+" "+b+" "+c+" "+p+" "+m+" "+n+" ")
                if (triangles[p] == 1) result++;
                if (triangles[p] == 2) result--;
                p += a+b+c;
            }
        }
    }
  }

  console.log("Answer is "+result);
  return result;

}

singularIntRightTriangles();
