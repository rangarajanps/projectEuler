var p = [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ];

function GetNextPerm() {

  var N = p.length;
  var i = N - 1;

  while (p[i - 1] >= p[i]) {
    i--;
    if (i < 1) return false;
  }

  var j = N;
  while (p[j - 1] <= p[i - 1]) {
    j = j - 1;
  }

  // swap values at position i-1 and j-1
  swap(i - 1, j - 1);

  i++;
  j = N;

  while (i < j) {
    swap(i - 1, j - 1);
    i++;
    j--;
  }
  return true;
}

function swap(i, j) {
  var k = p[i];
  p[i] = p[j];
  p[j] = k;
}

function CheckResult() {
            
  if (p[1] == 10 ||
      p[2] == 10 ||
      p[4] == 10 ||
      p[6] == 10 ||
      p[8] == 10) return false;
            
  if (p[0] > p[3] ||
      p[0] > p[5] ||
      p[0] > p[7] ||
      p[0] > p[9] ) return false;
            
  if (p[0] + p[1] + p[2] != p[3] + p[2] + p[4]) return false;
  if (p[0] + p[1] + p[2] != p[5] + p[4] + p[6]) return false;
  if (p[0] + p[1] + p[2] != p[7] + p[6] + p[8]) return false;
  if (p[0] + p[1] + p[2] != p[9] + p[8] + p[1]) return false;
          
  return true;
}

function magic5GonRing() {
  
  var result ="";
  while (true) {
    if (!GetNextPerm()) break;
    if (CheckResult()) {
       var candidate = "" + p[0] + p[1] + p[2] + p[3] + p[2] + p[4] + p[5] + p[4] + p[6] + p[7] + p[6] + p[8] + p[9] + p[8] + p[1];
       if (candidate.localeCompare(result) > 0) result = candidate;
     }                 
  }

  return parseInt(result);

}

//console.log(magic5GonRing());
