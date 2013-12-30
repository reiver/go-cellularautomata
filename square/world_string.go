package square



import "bytes"



func (me *World) String() string {
	var buffer bytes.Buffer

	timeIndex := me.CurrentTimeIndex()

	for h:=0; h<me.height; h++ {
		for w:=0; w<me.width; w++ {

			index := me.xyToIndex(w,h)

			buffer.WriteRune(me.temporalBuffer[timeIndex][index].Value)
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

