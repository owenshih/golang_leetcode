package golang

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var points []string
	var ans = 0
	for _, v := range operations {
		fmt.Println(v)
		switch v {
		case "+":
			points = plusPoint(points)
		case "D":
			points = doublePoint(points)
		case "C":
			points = removePoint(points)
		default:
			points = append(points, v)
		}
		fmt.Println(points)
	}

	for _, v := range points {
		var p, _ = strconv.Atoi(v)
		ans += p
	}
	return ans
}

func removePoint(points []string) []string {
	points = points[:len(points)-1]
	return points
}

func plusPoint(points []string) []string {
	var lastPoint, _ = strconv.Atoi(points[len(points)-1])
	var last2Point, _ = strconv.Atoi(points[len(points)-2])
	var point = lastPoint + last2Point
	points = append(points, strconv.Itoa(point))
	return points
}

func doublePoint(points []string) []string {
	var lastPoint, _ = strconv.Atoi(points[len(points)-1])
	var point = lastPoint * 2
	points = append(points, strconv.Itoa(point))
	return points
}
