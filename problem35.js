function isPrime(number){
	if(number <= 1){return false;}
	if(number == 2){return true;}
	if(!(number%2)){return false;}
	var maxCheck = Math.floor(Math.sqrt(number));
	for(var i = 3; i <= maxCheck; i+=2){
		if(!(number%i)){return false;}
	}
	return true;
}

function circularPrimes(n) {
  var numCircular = 0;  
  for(var i = 2; i < n; i++){
	  if (isPrime(i))  {
		  var circular = true;
		  var strI = String(i);
		  var length = strI.length;
		  for(var j = 0; j<length-1; j++){
			  var rotation = strI.substr(1)+strI[0];
			  strI = rotation;
			  if(!isPrime(parseInt(rotation))){circular=false;break;}
		  }
		  if  (circular) {
			  //console.log("Found a circular prime "+i);
			  numCircular++;
		  }
	  }
  }
  return numCircular;
}

console.log(circularPrimes(1000));
