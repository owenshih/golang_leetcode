package golang

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	min, max, guessNum := 1, n, (1+n)/2
	for {
		guessRes := guess(guessNum)
		if guessRes == 0 {
			return guessNum
		}
		if guessRes == 1 {
			min = guessNum + 1
		}
		if guessRes == -1 {
			max = guessNum - 1
		}
		guessNum = (min + max) / 2
	}

}
