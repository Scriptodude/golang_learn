// Few basic numbers computation algorithms
// Such as fibonacci sequence, factorial, etc.
package main

import ("errors";)

const MAX_UINT64 uint64 = ^uint64(0)
const QUARTER_MAX_UINT64 = (MAX_UINT64 >> 2)
const MAX_INT64 int64 = int64(MAX_UINT64 >> 1)
const QUARTER_MAX_INT64 int64 = (MAX_INT64 >> 2)
const QUARTER_MIN_INT64 int64 = -QUARTER_MAX_INT64

func fibonacci(n int) (int64, error) {
	if (n <= 2) {
		return 1, nil
	}

	prev1 := int64(0)
	var total int64 = 1
	for i := 2; i <= n; i++ {

		// If we're close to an overflow
		if (total >= QUARTER_MAX_INT64) {

			// We make sure with our n-1 and n-2 we don't overflow the integer
			diff := MAX_INT64 - total
			if(diff < int64(prev1)){
				return 0, errors.New("numbers : fibonacci integer 64 overflow.")
			}
		}

		temp := total
		total += prev1
		prev1 = temp
	}

	return total, nil
}

func factorial(n uint64) (uint64, error) {
	// If n is really large, it is obvious that it will overflow.
	if (n > QUARTER_MAX_UINT64){
		return 0, errors.New("numbers : factorial unsigned interger 64 overflow.")
	}

	var total uint64 = 1 
	for ; n > 0; n-- {

		// Case we're going to overflow
		if(total * 2 > QUARTER_MAX_UINT64){
			return 0, errors.New("numbers : factorial unsigned interger 64 overflow.")
		}

		total *= n
	}

	return total, nil
}