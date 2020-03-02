function getDigit(n) {
 var r, s = 0, k = 0;
 while (s < n) {
   r = s;
   s+= 9 * 10**k++ * k;
 }
 //console.log("r s k is "+r+" "+s+" "+k);
 var h = n - r - 1;
 var t = 10**(k - 1) + h / k;
 var p = h % k;
 //console.log("h t p is "+h+" "+t+" "+p);
 return +String(t)[p];
}

function champernownesConstant(n) {
  var res = 1;
  var max = String(n).length;
  for (var i = 0; i < max; i++) {
    res*= getDigit(10**i);
  }
  return res;
}

console.log(champernownesConstant(100));
