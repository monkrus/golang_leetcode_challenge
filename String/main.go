1. Reverse string
Write a function that reverses a string. The input string is given as an array of characters char[].


//  it takes string and converts to bytes
func reverseString(s []byte)  {

	// i is a starting index
	// j is the last index of array
		
	i := 0
	j := len(s) - 1
	
	// until i=j 
	for i < j {
	// first index = last index, last index = first index
	s[i], s[j] = s[j], s[i]
	
	// add 1 from the start
	i++
	//substract one from the last index
	j--
	   
		}
	}
	