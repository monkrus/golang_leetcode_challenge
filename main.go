1. Remove duplicates from Sorted array

// we set a function that take nums array, returns int.
// If the length of array is less than two, which means there are NO duplicates,
// it returns the length of it (ex. 0 or 1 )

func removeDuplicates(nums []int) int {
 if len(nums) < 2 {
		return len(nums)
	}
 
 /*we set starting element (one item in array) and preVal (the value of the first element in array)
   run the for loop from 1 to the whole length of the array
   if the length of the nums is NOT equal to preVal (ex. there is more than one element in array)
 
    start gets +1 element ()
    preval gets a value of the new element
    and nums[start] becomes nums[i]
    (ex. start=1,2 ; preval=nums [0 1])
    nums[start] now equal to the nums[i]
    meaning the length of array has been adjusted
  */
 
	start, preVal := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			start, preVal, nums[start] = start+1, nums[i], nums[i]
		}
	}
	return start
   
}

2. Best time to buy and sell stock II

func maxProfit(prices []int) int {
   // initially the profit is zero
    profit := 0
   // we run a loop, where i is less the the lengths of the `prices`
   // therefore finding the max price
   // we get an array of numbers 
   // and find the pair of two prices (high + high minus one)
   // the difference between two prics will give us a the profit
   for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            profit += prices[i] - prices[i-1]
        }
    }
    
	return profit

}

