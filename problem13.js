function setCharAt(str,index,chr) {
    if(index > str.length-1) return str;
    return str.substr(0,index) + chr + str.substr(index+1);
}

function largeSum(arr) {
  
  var carryFwd = 0;
  var lastCarryFwd = 0;
  var ans = '0'.repeat(arr[0].length);
  
  for ( var i=0; i< arr.length; i++) {
    for ( var j=arr[0].length-1; j >= 0; j--) {
      let temp = carryFwd+parseInt(arr[i][j])+parseInt(ans[j]);
      carryFwd = Math.floor(temp / 10);
      ans = setCharAt(ans, j, (temp % 10));
    }
    lastCarryFwd = lastCarryFwd + carryFwd;
  }
  if ( lastCarryFwd != 0 ) {
    ans = lastCarryFwd+ans;
  }
  
  //console.log("Answer is "+ans);
  return parseInt(ans.substr(0,10));
}

// only change code above this line

const testNums = [
  '37107287533902102798797998220837590246510135740250',
  '46376937677490009712648124896970078050417018260538'
];

largeSum(testNums)
