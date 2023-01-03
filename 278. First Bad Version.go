package golang

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	min, max, guessNum := 1, n, (1+n)/2
	for {
		if min == max {
			return max
		}
		if min+1 == max {
			isBad := isBadVersion(min)
			if isBad {
				return min
			} else {
				return max
			}
		}
		isBad := isBadVersion(guessNum)
		if isBad {
			max = guessNum
		} else {
			min = guessNum
		}
		guessNum = (min + max) / 2
	}
}
