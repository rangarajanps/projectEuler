function isRoyalFlush(cards) {
  var charCards = "TJQKA";
  for ( var i = 0; i < charCards.length; i++) {

  }
}

function calculatePoints(cards) {

  var card = cards.split(" ");

  var suiteTypes = "HDSC";
  var charCards = "TJQKA";
  var charPoints = [10, 11, 12, 13, 14]; 
  // T 10 J 11 Q 12 K 13 A 14

  var cardSuite ="";
  var cardVal = "";
  for ( var i = 0; i < card.length; i++) {
    cardSuite = cardSuite.concat(card[i][1]);
    cardVal = cardVal.concat(card[i][0]);
  }
  console.log("Card values => "+cardVal+" Card suite => "+cardSuite);

  var royal = 0;
  for ( var i = 0; i < charCards.length; i++ ) {
    if ( cardVal.indexOf(charCards[i]) != -1 ) {
      royal++;
    }
  }
  if ( royal == charCards.length ) {
    console.log("Found Ten, Jack, Queen, King, Ace ");
  }

  var suiteCount = new Map();
  
  for ( var i = 0; i < suiteTypes.length; i++ ) {
    if ( cardSuite.indexOf(suiteTypes[i]) != -1 ) {
      suiteCount.set(suiteTypes[i], suiteCount.get(suiteTypes[i])+1)
    }
  }
  console.log("Suite count is ");
  suiteCount.forEach(function (key, value) {
    console.log("key "+key+" value "+value);
  });

  var points = 0;
  var temp = 0;

  for ( var i = 0; i < card.length; i++){
    //console.log("Card value is "+card[i][0]);
    if ( Number.isInteger(parseInt(card[i][0])) ) {
      points = points+parseInt(card[i][0]);
    } else {
      temp = charCards.indexOf(card[i][0]);
      //console.log("Char cards is "+temp);
      points = points+charPoints[temp];
    }
    //console.log("Calculated point is "+points);
  }

  return points;
}

function pokerHands(arr) {
  
  var player1Hand;
  var player2Hand;
  var player1Points = 0;
  var player2Points = 0;
  var player1WinCount = 0;
  var player2WinCount = 0;
  
  for ( var i = 0; i < arr.length; i++) {
    player1Hand = arr[i].slice(0,14);
    //console.log("Player 1 Hand is "+player1Hand);
    player2Hand = arr[i].slice(15);
    //console.log("Player 2 Hand is "+player2Hand);
    player1Points = calculatePoints(player1Hand);
    player2Points = calculatePoints(player2Hand);
    console.log("Player 1 points => "+player1Points);
    console.log("Player 2 points => "+player2Points);
    if ( player1Points > player2Points ) {
      player1WinCount++;
      console.log("Player 1 Wins => "+player1WinCount);
    } else {
      player2WinCount++;
      console.log("Player 2 Wins => "+player2WinCount);
    }
  }

  return player1WinCount;
}

const test1Arr = [
  '8C TS KC 9H 4S 7D 2S 5D 3S AC']

const testArr = [
  '8C TS KC 9H 4S 7D 2S 5D 3S AC',
  '5C AD 5D AC 9C 7C 5H 8D TD KS',
  '3H 7H 6S KC JS QH TD JC 2D 8S',
  'TH 8H 5C QS TC 9H 4D JC KS JS',
  '7C 5H KC QH JD AS KH 4C AD 4S'
];

console.log(pokerHands(test1Arr));
