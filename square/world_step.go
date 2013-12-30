package square


type CellGetterFunc func(x, y int) rune
type CellSetterFunc func(x, y int, r rune)
type CellStepFunc   func(x, y int, getCurrentCellValue CellGetterFunc, setNextCellValue CellSetterFunc)


func (me *World) SetCellStepFunc(cellStepFunc CellStepFunc) {

	me.cellStepFunc = cellStepFunc
}
func (me *World) CellStepFunc() (CellStepFunc) {

	return me.cellStepFunc
}



func (me *World) Step() {

	// For each cell on in the world.
		lenBuffer := me.width*me.height

		for i:=0; i<lenBuffer; i++ {

			// Convert the buffer index (i) to x-y-coordinates.
				x,y := me.indexToXY(i)


			// Invoke the Cell Step Func.
				me.cellStepFunc(x,y, me.Get, me.setNext)

		} // for


	// Move the (new) calculated values in the "Next" field of the cell to the "Current" field.
		me.IncrementTimeIndex()
}
