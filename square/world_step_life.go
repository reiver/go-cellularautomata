package square


func LifeCellStepFunc(x, y int, getCurrentCellValue CellGetterFunc, setNextCellValue CellSetterFunc) {

	// Figure out how many neighbords the cell has.
		crowdSize := 0

		if ' ' != getCurrentCellValue(x-1 , y-1) {
			crowdSize++
		}
		if ' ' != getCurrentCellValue(x   , y-1) {
			crowdSize++
		}
		if ' ' != getCurrentCellValue(x+1 , y-1) {
			crowdSize++
		}

		if ' ' != getCurrentCellValue(x-1 , y) {
			crowdSize++
		}
		if ' ' != getCurrentCellValue(x+1 , y) {
			crowdSize++
		}

		if ' ' != getCurrentCellValue(x-1 , y+1) {
			crowdSize++
		}
		if ' ' != getCurrentCellValue(x   , y+1) {
			crowdSize++
		}
		if ' ' != getCurrentCellValue(x+1 , y+1) {
			crowdSize++
		}


	// If the cell is dead but the cell has (exactly) 3 neighbords, then make the cell live.
	// Else (if the cell is alive then) if the live cell has either 2 or 3 neighbords, then it stays alive, else make it dead.
	//
	// We store these (new) calculated values in the "Next" field in the cell, because we need the (old) "Current" values to stay
	// there until we have completed all the calculations.
		if ' ' == getCurrentCellValue(x,y) {

			if 3 == crowdSize {
				setNextCellValue(x,y, '\u2593')
			} else {
				setNextCellValue(x,y, ' ')
			}

		} else {

			if 2 > crowdSize || 3 < crowdSize {
				setNextCellValue(x,y, ' ')
			} else {
				setNextCellValue(x,y, '\u2588')
			}
		}
}
