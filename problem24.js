function lexicographicPermutations(n) {
    
    //Factorial lookup
		var factorials = [1,1,1,1,1,1,1,1,1,1];
		for(var i = 1; i<10; i++){
			factorials[i] = i*factorials[i-1];
		}

		var permutation = [0,1,2,3,4,5,6,7,8,9];
		var nth = 1;
		var max = n+1;
		var level = 9;
		var higher,current;
		while ( nth < max && level ) {

			if ( max - nth >= factorials[level] ) {
				//Increasing nth by level!
				nth += factorials[level];

				//Get next highest digit
				higher = 9-level+1;
				current = 9-level;
				while(permutation[higher] < permutation[current]){higher++;}

				//Swap current and next highest
				var tmp = permutation[current];
				permutation[current]  = permutation[higher];
				permutation[higher] = tmp;

			} else {	
				level--;
			}
		}

		return permutation.join("");
}

//console.log(lexicographicPermutations(999999));
