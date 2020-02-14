function sumOfWord(word) {

  var sum = 0;
  for (var i = 0; i < word.length; i++) {
    sum = sum + word.charCodeAt(i)-64;
    //console.log("Sum is "+sum+" for i "+i+" char "+word[i]);
  }
  //console.log("Sum for word "+word+" is "+sum);
  return sum;
}

function namesScores(arr) {
  
  var sortedArr = arr.sort();
  var tempScore = 0;
  var totalScore = 0;
  for (var i = 0; i < sortedArr.length; i++) {
    tempScore = tempScore + sumOfWord(sortedArr[i]);
    tempScore = tempScore * (arr.indexOf(sortedArr[i])+1);
    //console.log("tempScore is "+tempScore+" for i "+i+" word "+sortedArr[i]);
    totalScore = totalScore + tempScore;
    tempScore = 0;
  }

  return totalScore;

}

// Only change code above this line
const test1 = ['THIS', 'IS', 'ONLY', 'A', 'TEST'];
const test2 = ['I', 'REPEAT', 'THIS', 'IS', 'ONLY', 'A', 'TEST'];

//console.log(namesScores(test1));
namesScores(test1);
