package square


func (me *World) TimeIndexInit() {
	me.timeIndex = 0
}

func (me *World) CurrentTimeIndex() uint {
	return me.timeIndex
}

func (me *World) NextTimeIndex() uint {
	// Increment
		nextTimeIndex := me.CurrentTimeIndex() + 1


	// Canonicalize
		nextTimeIndex = nextTimeIndex % uint(len(me.temporalBuffer))


	// Return
		return nextTimeIndex
}

func (me *World) IncrementTimeIndex() {
	me.timeIndex = me.NextTimeIndex()
}
