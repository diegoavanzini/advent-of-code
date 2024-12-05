package day4

func ceresSearchMasX(input []string) (found int) {
	aLocations := searchFor('A', input)
	for i := 0; i < len(aLocations); i++ {
		if aLocations[i].IsMas(input) {
			found++
		}
	}
	return
}

func ceresSearch(word string, input []string) (found int) {
	var locations = make([][]Location, len(word))
	for i := 0; i < len(word); i++ {
		locations[i] = searchFor(word[i], input)
	}
	return searchWord(locations)

}

func searchWord(locations [][]Location) (found int) {
	xLocations := locations[0]
	mLocations := locations[1]

	for iX := 0; iX < len(xLocations); iX++ {
		for iM := 0; iM < len(mLocations); iM++ {
			if dir, ok := xLocations[iX].Near(mLocations[iM]); ok {
				found += searchNear(locations[2:], mLocations[iM], dir)
			}
		}
	}
	return
}

func searchNear(currentCharLocations [][]Location, current Location, dir Direction) (found int) {
	if len(currentCharLocations) == 0 {
		return 1
	}
	for i := 0; i < len(currentCharLocations[0]); i++ {
		if newDir, ok := current.Near(currentCharLocations[0][i]); ok &&
			newDir == dir {
			found += searchNear(currentCharLocations[1:], currentCharLocations[0][i], dir)
			return
		}
	}
	return
}

func searchFor(b byte, input []string) (xLocations []Location) {
	for iRow := 0; iRow < len(input); iRow++ {
		for iCol := 0; iCol < len(input[iRow]); iCol++ {
			if input[iRow][iCol] == b {
				xLocations = append(xLocations, Location{Row: iRow, Col: iCol})
			}
		}
	}
	return
}
