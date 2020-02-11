var daysInMonth = [31,28,31,30,31,30,31,31,30,31,30,31]
var monthNames = ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"]
var dayInWeek = ["Su","M","Tu","W","Th","F","Sa"]

function isLeap(year) {
  if ( year % 4 == 0 ) {
    if ( year % 100 !=0 ) {
      //console.log(year+" is a Leap year");
      return true;
    }
    if ( year % 400==0 ) {
      //console.log(year+" is a Leap year");
      return true;
    }
  } 
  return false;
}

function findDayOf1stJanOfAYear(firstYear) {
  
  var day = 0; //1899 Jan 1st is Sunday
  for (var i = 1900; i <= firstYear; i++) {
    day = day + 1;
    if ( isLeap(i)) {
      day = day + 1;
    } 
  }
  
  return (day%7);
}

function countingSundays(firstYear, lastYear) {
  
  var day = 0;
  var countOfSundays = 0;
  var firstOfAYear = findDayOf1stJanOfAYear(firstYear);

  for (var i = firstYear; i <= lastYear; i++) {

    day = day + firstOfAYear;
    
    for (var j = 0; j < daysInMonth.length; j++) {

      if ( day % 7 == 0 ) {
        //console.log("First of "+monthNames[j]+" month in the year "+i+" is a Sunday");
        countOfSundays++;
      }
      day = day + daysInMonth[j];
      if ( j == 1 && isLeap(i) ) {
        //console.log("Incrementing a day for leap year "+i)
        day = day + 1;
      }
    }

    day = 0;
    firstOfAYear = firstOfAYear + 1;
    if ( isLeap(i)) {
      firstOfAYear = firstOfAYear + 1;
    }
  }

  return countOfSundays;
}

//console.log("Number of Sundays is "+countingSundays(1943, 1946));
countingSundays(1943, 1946);
