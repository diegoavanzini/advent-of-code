package historianhysteria

import (
	"math"
	"slices"
)

func totalDistance(leftList []int, rightList []int) (totalDistance int) {
	// simplest solution
	slices.Sort(leftList)
	slices.Sort(rightList)
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return totalDistance
}

func similarityScore(leftList []int, rightList []int) (similarityScore int) {
	for iLeft := 0; iLeft < len(leftList); iLeft++ {
		for iRight := 0; iRight < len(rightList); iRight++ {
			if leftList[iLeft] == rightList[iRight] {
				similarityScore += leftList[iLeft]
			}
		}
	}
	return
}
