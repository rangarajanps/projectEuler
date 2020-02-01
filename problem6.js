function sumSquareDifference(n) {
  var value1 = 0;
  var value2 = 0;
  for (var i = 1; i <= n; i++) {
    value1 = value1 + i*i;
    value2 = value2 + i;
  }

  value2 = value2 * value2;

  return value2-value1;
}

sumSquareDifference(100);
