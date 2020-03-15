1. Remove duplicates from Sorted array
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

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

2. Best time to buy and sell stock II.

Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit.
You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

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

3. Rotate array
Given an array, rotate the array to the right by k steps, where k is non-negative.

//  if k is 0, rotation does not make sense
//  if k is equal to the length of array, it will return to the initial position
//  if k is more than a length, k equals to k moduls of length (e.g. len is 5, k is 7, the modulus is 2)


func rotate(nums []int, k int)  {
 if k == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}
// we reverse the array nums, cutting the slice between 0th element and the length of the array `nums` minus k
// then we reverse the just created slice
	reverse(nums[0 : len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

//

func reverse(s []int) {
	if len(s) == 0 {
		return
	}
	end := len(s) - 1
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[end-i] = s[end-i], s[i]
	}   
}


