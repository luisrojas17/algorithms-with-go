package algorithms

import "fmt"

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which
// are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need
// to do the following things:
//
// Change the array nums such that the first k elements of nums contain the elements which are not equal
// to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.
//
// Custom Judge:
//
// The judge will test your solution with the following code:
//
// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
// It is sorted with no values equaling val.
//
// int k = removeElement(nums, val); // Calls your implementation
//
// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
// for (int i = 0; i < actualLength; i++) {
//     assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.
//
// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
//It does not matter what you leave beyond the returned k (hence they are underscores).
//
// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).
//
// Constraints:
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100
func InitRemoveZeroElements() {

	fmt.Println("Executing Remove Element Algorithm")

	// Case 1
	numbers := []int{3, 2, 2, 3}
	val := 3
	result := removeV2(numbers, val)
	fmt.Println("\tResult for version 2: ", numbers[0:result])

	//Case 2
	numbers = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2

	result = removev1(numbers, val)
	fmt.Println("\tResult for version 1: ", numbers[0:result])
}

func removev1(nums []int, val int) int {

	// This variable specify the movements
	var result int
	if nil != nums {

		// [0, 1, 2, 2, 3, 0, 4, 2] -> [0,1,4,0,3,_,_,_]
		//  0, 1, 2, 3, 4, 5, 6, 7
		//  i, j
		for i := 0; i < len(nums); i++ {

			if nums[i] != val {

				nums[result] = nums[i]

				// Set numer of movements and next index where the value will be set.
				result++
			}

		}

	}

	return result
}

func removeV2(nums []int, val int) int {

	var tempValue = 0
	var fakeValue = -1
	var result = len(nums)

	for i := 0; i < len(nums); i++ {

		// To chnage the value for default value to simulate remove it
		if nums[i] == val {

			// [0, 1, 2, 2, 3, 0, 4, 2] -> [0,1,4,2,3,0,2,2]
			//        i
			// [0, 1, 4, 2, 3, 0, 2, 2] -> [0,1,4,0,3,0,2,2]
			//           i
			nums[i] = fakeValue

			// To indicate the numbers that are not equals to val.
			result--

			// To iterate the array in reverse order and moving the current item to
			// last position because the current item corresponds with value to remove
			for j := len(nums) - 1; j > i; j-- {

				// [0, 1, 2, 2, 3, 0, 4, 2] -> [0,1,4,2,3,0,2,2]
				//        i           j
				// [0, 1, 4, 2, 3, 0, 2, 2] -> [0,1,4,0,3,2,2,2]
				//           i     j
				// If last item is different to val and fakeValue
				// it will be changing the positions.
				if nums[j] != val && nums[j] != fakeValue {

					// Save the last item's value.
					tempValue = nums[j]

					// Move the item's value to remove at the end of the array.
					nums[j] = nums[i]

					// Move the last item's value at the begining of the array.
					nums[i] = tempValue

					break
				}
			}
		}

	}

	return result
}
