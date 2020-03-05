function isPentagonal(n) {
  let p = (1 + Math.sqrt(1 + 24*n)) / 6
  return p === Math.floor(p)
}

function solvePentagonal(n) {
  return (n*(3*n - 1))/2;
}

function pentagonNumbers() {
  let found = false;
  let i = 1;
  let j, p1, p2, ans;

  while (!found) {
    p1 = solvePentagonal(i);
    for (j = i - 1; j > 0; j--) {
      p2 = solvePentagonal(j);
      if (isPentagonal(p1 + p2) && isPentagonal(p1 - p2)) {
        found = true;
        ans = p1 - p2;
        break;
      }
    }
    i++;
  }

  return(ans);
}

pentagonNumbers();
