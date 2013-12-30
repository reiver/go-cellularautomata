package square



import (
	"math/rand"
	"testing"
	"time"
)



func TestTime(t *testing.T) {

	// Create.
		randomness := rand.New(rand.NewSource( time.Now().UnixNano() ))

		worldWidth  := 1 + randomness.Intn(20)
		worldHeight := 1 + randomness.Intn(20)

		world := NewWorld(worldWidth, worldHeight)



	// Confirm widith
		actualWidth := world.Width()

		if worldWidth != actualWidth {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected width of [%v].", worldWidth, worldHeight, actualWidth)
			return
		}



	// Confirm widith
		actualHeight := world.Height()

		if worldHeight != actualHeight {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected height of [%v].", worldWidth, worldHeight, actualHeight)
			return
		}



	// Set the current time index to something else (so that when we reset it, it will actually change).
		otherTimeIndex := uint(13)
		world.timeIndex = otherTimeIndex

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; otherTimeIndex != actualCurrentTimeIndex {
			t.Errorf("Tried changing current time index to %v but when checked, it was actually %v.", otherTimeIndex, actualCurrentTimeIndex)
			return
		}



	// Confirm init works
		world.TimeIndexInit()

		expectedCurrentTimeIndex := uint(0)

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; expectedCurrentTimeIndex != actualCurrentTimeIndex {
			t.Errorf("After init'ing, got current time index and expected it to be %v, but when checked, it was actually %v.", expectedCurrentTimeIndex, actualCurrentTimeIndex)
			return
		}

		expectedNextTimeIndex := uint(1)

		if actualNextTimeIndex := world.NextTimeIndex() ; expectedNextTimeIndex != actualNextTimeIndex {
			t.Errorf("After init'ing, got next time index and expected it to be %v, but when checked, it was actually %v.", expectedNextTimeIndex, actualNextTimeIndex)
			return
		}


	// Confirm increment works #1
		world.IncrementTimeIndex()

		expectedCurrentTimeIndex = 1

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; expectedCurrentTimeIndex != actualCurrentTimeIndex {
			t.Errorf("After incrementing once, got current time index and expected it to be %v, but when checked, it was actually %v.", expectedCurrentTimeIndex, actualCurrentTimeIndex)
			return
		}

		expectedNextTimeIndex = 0

		if actualNextTimeIndex := world.NextTimeIndex() ; expectedNextTimeIndex != actualNextTimeIndex {
			t.Errorf("After incrementing once, got next time index and expected it to be %v, but when checked, it was actually %v.", expectedNextTimeIndex, actualNextTimeIndex)
			return
		}


	// Confirm increment works #2
		world.IncrementTimeIndex()

		expectedCurrentTimeIndex = 0

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; expectedCurrentTimeIndex != actualCurrentTimeIndex {
			t.Errorf("After incrementing twice, got current time index and expected it to be %v, but when checked, it was actually %v.", expectedCurrentTimeIndex, actualCurrentTimeIndex)
			return
		}

		expectedNextTimeIndex = 1

		if actualNextTimeIndex := world.NextTimeIndex() ; expectedNextTimeIndex != actualNextTimeIndex {
			t.Errorf("After incrementing twice, got next time index and expected it to be %v, but when checked, it was actually %v.", expectedNextTimeIndex, actualNextTimeIndex)
			return
		}


	// Confirm increment works #3
		world.IncrementTimeIndex()

		expectedCurrentTimeIndex = 1

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; expectedCurrentTimeIndex != actualCurrentTimeIndex {
			t.Errorf("After incrementing thrice, got current time index and expected it to be %v, but when checked, it was actually %v.", expectedCurrentTimeIndex, actualCurrentTimeIndex)
			return
		}

		expectedNextTimeIndex = 0

		if actualNextTimeIndex := world.NextTimeIndex() ; expectedNextTimeIndex != actualNextTimeIndex {
			t.Errorf("After incrementing thrice, got next time index and expected it to be %v, but when checked, it was actually %v.", expectedNextTimeIndex, actualNextTimeIndex)
			return
		}


	// Confirm increment works #4
		world.IncrementTimeIndex()

		expectedCurrentTimeIndex = 0

		if actualCurrentTimeIndex := world.CurrentTimeIndex() ; expectedCurrentTimeIndex != actualCurrentTimeIndex {
			t.Errorf("After incrementing fource, got current time index and expected it to be %v, but when checked, it was actually %v.", expectedCurrentTimeIndex, actualCurrentTimeIndex)
			return
		}

		expectedNextTimeIndex = 1

		if actualNextTimeIndex := world.NextTimeIndex() ; expectedNextTimeIndex != actualNextTimeIndex {
			t.Errorf("After incrementing fource, got next time index and expected it to be %v, but when checked, it was actually %v.", expectedNextTimeIndex, actualNextTimeIndex)
			return
		}
}
