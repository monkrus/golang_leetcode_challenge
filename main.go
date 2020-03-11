1. Remove duplicates from Sorted array

// we set a function that take nums array, returns int.
// If the length of array is less than two,which means there are no duplicates,
// it returns the lenght of it (ex. 0 or 1 )

func removeDuplicates(nums []int) int {
 if len(nums) < 2 {
		return len(nums)
	}
 
 //we set start (one item in array) and preVal (first element in array)
 // run the for loop from 1 to the whole length of the array
 // if the length of the nums is NOT equal to preVal (ex. there is more than one element in array)
 
 // start gets +1 value ( until it will be equal to nums[]
 // preval gets a value of nums[i], and nums[start] becomes nums[1]
 // (ex. start=2, preval=nums [0 1], nums[start] becomes a nums[i]
 
	start, preVal := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			start, preVal, nums[start] = start+1, nums[i], nums[i]
		}
	}
	return start
   
}
