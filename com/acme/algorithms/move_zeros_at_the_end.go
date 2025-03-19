package algorithms

import "fmt"

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
// Input: nums = [0]
// Output: [0]
//
// Constraints:
//
// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1
func InitMoveZerosAtTheEnd() {

	fmt.Println("Executing Move Zeros at the end of the Array algorithm...")

	var numbers = []int{0, 1, 0, 3, 12}

	moveZerosV1(numbers)

	fmt.Println("\t", numbers)

}

const ZERO = 0

func moveZerosV1(nums []int) {

	if nums != nil {

		// [0, 1, 0, 3, 12]
		for i := 0; i < len(nums); i++ {

			// If the current value = 0, it will be move to next position.
			if nums[i] == ZERO {

				// [0, 1, 0, 3, 12] -> // [1, 0, 0, 3, 12]
				//  i, j, j, j, j
				// Start 'j' to next position according to 'i' position
				for j := i + 1; j < len(nums); j++ {

					// Change the position of i and j values
					if nums[j] != ZERO {

						tempValue := nums[j]

						nums[j] = nums[i]

						nums[i] = tempValue

						break
					}

				}

			}

		}

	}

}
