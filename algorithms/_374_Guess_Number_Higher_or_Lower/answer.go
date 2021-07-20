package _374_Guess_Number_Higher_or_Lower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	var (
		low  = 1
		high = n
	)
	for low <= high {
		mid := (low + high) / 2
		answer := guess(mid)
		if answer == 0 {
			return mid
		} else if answer == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
