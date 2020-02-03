function specialPythagoreanTriplet(n) {
 
 var a = 1;
 var b = a+1;
 var c = b+1; 
 while (a+b+c <= n) {

     for ( var x = a+1; a+x+c <= n; x++) {
        c = Math.sqrt(a*a + x*x);
        if ( (a+x+c) == n ) {
            return a*x*c;
        }
     }
     
     a++;
     b = a+1;
     c = b+1;

 }

 return false;
}

specialPythagoreanTriplet(12);
