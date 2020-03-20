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
func rotate(nums []int, k int)  {
    n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

4. Contains duplicate

/*Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.*/

// func takes an array nums, returns bool
// record is a map with int key/ int value
// for loop an array , the range 
func containsDuplicate(nums []int) bool {
	record := make(map[int]int)
        for _, num := range nums {
	if _, ok := record[num]; 
	!ok {
	record[num] = 1
	} else {
	return true
	}
	}
	return false
        }


	5. Single numbers

	func singleNumber(nums []int) int {
		res := 0
		for _, n := range nums {
			// n^n == 0
			// a^b^a^b^a == a
			res ^= n
		}
		return res
	}

	/* ^ A bitwise XOR takes two bit patterns of equal length and performs the logical exclusive OR operation on each pair of corresponding bits. 
	The result in each position is 1 if only the first bit is 1 or only the second bit is 1, but will be 0 if both are 0 or both are 1. 
	In this we perform the comparison of two bits, being 1 if the two bits are different, and 0 if they are the same  

        1 -> 0/1 or 1/0
        0 -> 1/1 or 0/1

        if res = n, 0
        if res !=n, 1
        */
