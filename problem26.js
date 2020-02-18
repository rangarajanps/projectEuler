function resetCounterNumerator() {
  counter = 0;
  numerator = 1;
}

function reciprocalCycles(n) {
  
  var rem = 0;
  var max_denominator = 0;

  for ( var denominator = 2; denominator <= n; denominator++ ) {
    
    var init_rem = numerator % denominator;
    counter++;
    numerator = init_rem * 10;
    rem = init_rem;

    while ( counter < n ) {

      numerator = rem * 10;      
      rem = numerator % denominator;
      counter++;

      if ( rem == 0 ) {
        //console.log("Reminder is ZERO for "+denominator+" recurring "+counter)
        resetCounterNumerator();
        break;
      }

      if ( rem == init_rem ) {
        if ( counter >= n ) {
          //console.log("Answer found denominator "+denominator+" recurring "+counter);
          return denominator;
        }
        
        //console.log("Recurrence while detected for denominator "+denominator+" recurring "+counter);
        if ( counter > max_denominator ) {
          max_denominator = denominator;
        }
        
        resetCounterNumerator();
        break;
      }
      
      if ( counter >= n ) {
          //console.log("Recurrence NOT detected for denominator "+denominator+" recurring "+counter);
          resetCounterNumerator();
          break;
      }
    }

  }
  return max_denominator;
}

var numerator = 1;
var counter = 0;
console.log(reciprocalCycles(800));
