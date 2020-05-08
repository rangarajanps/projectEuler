function substringDivisibility() {
  var nums = [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 ];
	var result = [];

	// Run the algorithm until count is equal to n
	while (!inReverseOrder(nums)) {
		var i = nums.length - 1;
		while (nums[i - 1] >= nums[i])
		  i = i - 1;

		var j = nums.length;
		while (nums[j - 1] <= nums[i - 1])
			j = j - 1;

		swap(nums, i - 1, j - 1);

		i++;
		j = nums.length;
		while (i < j) {
			swap(nums, i - 1, j - 1);
			i++;
			j--;
		}

		if (isSubDivisible(nums)) {
			result.push(parseInt(nums.join("")));
		}
	}

	return result;

}

function swap(arr, i, j) {
	var temp = arr[i];
	arr[i] = arr[j];
	arr[j] = temp;
}

function inReverseOrder(arr) {
	for (var i = 0, j = arr.length - 1; i < arr.length; i++, j--) {
		if (arr[i] != j) {
			return false;
		}
	}
	return true;
}

function isSubDivisible(numbers) {

		// d2d3d4 = 406 is divisible by 2
		// d3d4d5 = 063 is divisible by 3
		// d4d5d6 = 635 is divisible by 5
		// d5d6d7 = 357 is divisible by 7
		// d6d7d8 = 572 is divisible by 11
		// d7d8d9 = 728 is divisible by 13
		// d8d9d10 = 289 is divisible by 17
		var d2d3d4 = numbers[1] * 100 + numbers[2] * 10 + numbers[3];
		if (d2d3d4 % 2 != 0) {
			return false;
		}
		var d3d4d5 = numbers[2] * 100 + numbers[3] * 10 + numbers[4];
		if (d3d4d5 % 3 != 0) {
			return false;
		}
		var d4d5d6 = numbers[3] * 100 + numbers[4] * 10 + numbers[5];
		if (d4d5d6 % 5 != 0) {
			return false;
		}
		var d5d6d7 = numbers[4] * 100 + numbers[5] * 10 + numbers[6];
		if (d5d6d7 % 7 != 0) {
			return false;
		}
		var d6d7d8 = numbers[5] * 100 + numbers[6] * 10 + numbers[7];
		if (d6d7d8 % 11 != 0) {
			return false;
		}
		var d7d8d9 = numbers[6] * 100 + numbers[7] * 10 + numbers[8];
		if (d7d8d9 % 13 != 0) {
			return false;
		}
		var d8d9d10 = numbers[7] * 100 + numbers[8] * 10 + numbers[9];
		if (d8d9d10 % 17 != 0) {
			return false;
		}

		return true;

}

//console.log(substringDivisibility());

//Output does not appear as per order expected by freecodecamp site
// running tests
//substringDivisibility() should return [ 1430952867, 1460357289, 1406357289, 4130952867, 4160357289, 4106357289 ].
// tests completed
// console output

//[ 1406357289,
//  1430952867,
//  1460357289,
//  4106357289,
//  4130952867,
//  4160357289 ]
