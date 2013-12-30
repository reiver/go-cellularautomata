package square



type World struct {
	timeIndex uint
	width  int
	height int
	cellStepFunc CellStepFunc
	temporalBuffer [][]Cell
}

func NewWorld(width int, height int) (*World) {

	// Calculate the length of a buffer.
		newBufferLen := width*height

	// Specify how many buffers we are going to have.
	// Must be at least 2 -- must be ≥ 2.
		historyDepth := 2


	// Create the buffers.
		newTemporalBuffer := make([][]Cell, historyDepth)

		for i:=0; i<historyDepth; i++ {

			newBuffer := make([]Cell, newBufferLen)

			for ii:=0; ii<len(newBuffer); ii++ {
				newBuffer[ii].Value = ' '
			}

			newTemporalBuffer[i] = newBuffer
		}


	// New.
		me := World{
			timeIndex      : uint(0),
			width          : width,
			height         : height,
			cellStepFunc   : LifeCellStepFunc,
			temporalBuffer : newTemporalBuffer,
		}

		me.TimeIndexInit()


	// Return.
		return &me
}



func (me *World) Width() int {
	return me.width
}
func (me *World) Height() int {
	return me.height
}



// Converts (x,y) into their canonical form.
// The canonical form is such that 0 ≤ x < me.width and 0 ≤ y < me.height;
// we use modular arithmetic to put it in that range.
func (me *World) canonicalXY(x, y int) (int, int) {

	xx := x % me.width
	yy := y % me.height

	if 0 > xx {
		xx = xx + me.width
	}
	if 0 > yy {
		yy = yy + me.height
	}

	return xx,yy
}



func (me *World) xyToIndex(x,y int) int {

	x,y = me.canonicalXY(x,y)

	index := y*me.width + x

	return index
}
func (me *World) indexToXY(index int) (x, y int) {
	x = index % me.width
	y = ((index - x) / me.width) % me.height

	// Just in case. (Shouldn't be needed though.)
	x,y = me.canonicalXY(x,y)

	return x,y
}



func (me *World) setNext(x, y int, value rune) {

	temporalIndex := me.NextTimeIndex()

	index := me.xyToIndex(x,y)

	me.temporalBuffer[temporalIndex][index].Value = value
}
func (me *World) Set(x, y int, value rune) {

	temporalIndex := me.CurrentTimeIndex()

	index := me.xyToIndex(x,y)

	me.temporalBuffer[temporalIndex][index].Value = value
}
func (me *World) Get(x, y int) rune {

	temporalIndex := me.CurrentTimeIndex()

	x,y = me.canonicalXY(x,y)

	index := me.xyToIndex(x,y)

	r := me.temporalBuffer[temporalIndex][index].Value

	return r
}
