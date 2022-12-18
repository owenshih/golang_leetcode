package golang

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	TotalDestination := 0
	preDest := 0
	firstIndex, lastIndex := start, destination
	if lastIndex < firstIndex {
		firstIndex, lastIndex = lastIndex, firstIndex
	}
	for i := range distance {
		if i < firstIndex || i >= lastIndex {
			preDest += distance[i]
		} else {
			TotalDestination += distance[i]
		}
	}
	if preDest < TotalDestination {
		return preDest
	}
	return TotalDestination
}
