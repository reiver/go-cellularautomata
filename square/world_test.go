package square



import (
	"testing"
)



func TestWidthAndHeight(t *testing.T) {

	tests := []struct {
		width  int
		height int
	}{
		{1 , 1},
		{2 , 2},
		{3 , 3},
		{4 , 4},
		{5 , 5},

		{2 , 1},
		{3 , 1},
		{4 , 1},
		{5 , 1},

		{1 , 2},
		{1 , 3},
		{1 , 4},
		{1 , 5},

		{2 , 7},
		{31 , 27},
	}


	for _,datum := range tests {

		// Create.
		world := NewWorld(datum.width, datum.height)



		actualWidth := world.Width()

		if datum.width != actualWidth {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected width of [%v].", datum.width, datum.height, actualWidth)
			continue
		}



		actualHeight := world.Height()

		if datum.height != actualHeight {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected height of [%v].", datum.width, datum.height, actualHeight)
			continue
		}
	}
}



func TestXYToIndex (t *testing.T) {

	tests := []struct {
		width, height int
		x,y int
		expectedIndex int
	}{
		{
			1,1,
			0,0,
			0,
		},



		{
			2,2,
			0,0,
			0,
		},

		{
			2,2,
			1,0,
			1,
		},

		{
			2,2,
			0,1,
			2,
		},

		{
			2,2,
			1,1,
			3,
		},

		{
			2,2,
			2,0,
			0,
		},

		{
			2,2,
			0,2,
			0,
		},



		{
			3,5,
			0,0,
			0,
		},

		{
			3,5,
			1,0,
			1,
		},

		{
			3,5,
			2,0,
			2,
		},

		{
			3,5,
			3,0,
			0,
		},

		{
			3,5,
			4,0,
			1,
		},

		{
			3,5,
			5,0,
			2,
		},

		{
			3,5,
			6,0,
			0,
		},

		{
			3,5,
			7,0,
			1,
		},



		{
			3,5,
			0,1,
			3,
		},

		{
			3,5,
			1,1,
			4,
		},

		{
			3,5,
			2,1,
			5,
		},

		{
			3,5,
			3,1,
			3,
		},

		{
			3,5,
			4,1,
			4,
		},

		{
			3,5,
			5,1,
			5,
		},

		{
			3,5,
			6,1,
			3,
		},

		{
			3,5,
			7,1,
			4,
		},



		{
			3,5,
			0,2,
			6,
		},

		{
			3,5,
			1,2,
			7,
		},

		{
			3,5,
			2,2,
			8,
		},

		{
			3,5,
			3,2,
			6,
		},

		{
			3,5,
			4,2,
			7,
		},

		{
			3,5,
			5,2,
			8,
		},

		{
			3,5,
			6,2,
			6,
		},

		{
			3,5,
			7,2,
			7,
		},



		{
			3,5,
			0,3,
			9,
		},

		{
			3,5,
			1,3,
			10,
		},

		{
			3,5,
			2,3,
			11,
		},

		{
			3,5,
			3,3,
			9,
		},

		{
			3,5,
			4,3,
			10,
		},

		{
			3,5,
			5,3,
			11,
		},

		{
			3,5,
			6,3,
			9,
		},

		{
			3,5,
			7,3,
			10,
		},



		{
			3,5,
			0,4,
			12,
		},

		{
			3,5,
			1,4,
			13,
		},

		{
			3,5,
			2,4,
			14,
		},

		{
			3,5,
			3,4,
			12,
		},

		{
			3,5,
			4,4,
			13,
		},

		{
			3,5,
			5,4,
			14,
		},

		{
			3,5,
			6,4,
			12,
		},

		{
			3,5,
			7,4,
			13,
		},



		{
			3,5,
			0,5,
			0,
		},

		{
			3,5,
			1,5,
			1,
		},

		{
			3,5,
			2,5,
			2,
		},

		{
			3,5,
			3,5,
			0,
		},

		{
			3,5,
			4,5,
			1,
		},

		{
			3,5,
			5,5,
			2,
		},

		{
			3,5,
			6,5,
			0,
		},

		{
			3,5,
			7,5,
			1,
		},



		{
			3,5,
			0,6,
			3,
		},

		{
			3,5,
			1,6,
			4,
		},

		{
			3,5,
			2,6,
			5,
		},

		{
			3,5,
			3,6,
			3,
		},

		{
			3,5,
			4,6,
			4,
		},

		{
			3,5,
			5,6,
			5,
		},

		{
			3,5,
			6,6,
			3,
		},

		{
			3,5,
			7,6,
			4,
		},
	}



	for _,datum := range tests {

		// Create.
		world := NewWorld(datum.width, datum.height)



		actualWidth := world.Width()

		if datum.width != actualWidth {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected width of [%v].", datum.width, datum.height, actualWidth)
			continue
		}



		actualHeight := world.Height()

		if datum.height != actualHeight {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected height of [%v].", datum.width, datum.height, actualHeight)
			continue
		}



		actualIndex := world.xyToIndex(datum.x, datum.y)

		if datum.expectedIndex != actualIndex {
			t.Errorf("For xyToIndex() on world with input width,height = %vx%v, for x,y = %d,%d expected index to be [%v] but instead got back [%v].", datum.width, datum.height, datum.x, datum.y, datum.expectedIndex, actualIndex)
			continue
		}		

	}
}



func TestIndexToXY (t *testing.T) {

	tests := []struct {
		width, height int
		index int
		expectedX,expectedY int
	}{
		{
			1,1,
			0,
			0,0,
		},



		{
			2,2,
			0,
			0,0,
		},

		{
			2,2,
			1,
			1,0,
		},

		{
			2,2,
			2,
			0,1,
		},

		{
			2,2,
			3,
			1,1,
		},



		{
			2,3,
			0,
			0,0,
		},

		{
			2,3,
			1,
			1,0,
		},

		{
			2,3,
			2,
			0,1,
		},

		{
			2,3,
			3,
			1,1,
		},

		{
			2,3,
			4,
			0,2,
		},

		{
			2,3,
			5,
			1,2,
		},



		{
			2,3,
			6,
			0,0,
		},

		{
			2,3,
			7,
			1,0,
		},

		{
			2,3,
			8,
			0,1,
		},

		{
			2,3,
			9,
			1,1,
		},

		{
			2,3,
			10,
			0,2,
		},

		{
			2,3,
			11,
			1,2,
		},

		{
			2,3,
			12,
			0,0,
		},

		{
			2,3,
			13,
			1,0,
		},
	}


	for _,datum := range tests {

		// Create.
		world := NewWorld(datum.width, datum.height)



		actualWidth := world.Width()

		if datum.width != actualWidth {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected width of [%v].", datum.width, datum.height, actualWidth)
			continue
		}



		actualHeight := world.Height()

		if datum.height != actualHeight {
			t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected height of [%v].", datum.width, datum.height, actualHeight)
			continue
		}



		actualX, actualY := world.indexToXY(datum.index)

		if datum.expectedX != actualX || datum.expectedY != actualY {
			t.Errorf("For indexToXY() on world with input width,height = %vx%v, for index = %d expteced x,y = %d,%d but instead got back %v,%v.", datum.width, datum.height, datum.index, datum.expectedX, datum.expectedY, actualX, actualY)
			continue
		}		

	}


}



func TestGetAndSet(t *testing.T) {

	tests := []struct {
		worldMap [][]rune
	}{
		{
			[][]rune{
				[]rune{' '},
			},
		},

		{
			[][]rune{
				[]rune{'#'},
			},
		},

		{
			[][]rune{
				[]rune{'a', 'b'},
				[]rune{'c', 'd'},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' '},
				[]rune{' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' '},
				[]rune{' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', '#'},
				[]rune{' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' '},
				[]rune{'#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' '},
				[]rune{' ', '#'},
			},
		},

		{
			[][]rune{
				[]rune{'#', '#'},
				[]rune{' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' '},
				[]rune{'#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' '},
				[]rune{' ', '#'},
			},
		},

		{
			[][]rune{
				[]rune{' ', '#'},
				[]rune{'#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', '#'},
				[]rune{' ', '#'},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' '},
				[]rune{'#', '#'},
			},
		},

		{
			[][]rune{
				[]rune{' ', '#'},
				[]rune{'#', '#'},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' '},
				[]rune{'#', '#'},
			},
		},

		{
			[][]rune{
				[]rune{'#', '#'},
				[]rune{' ', '#'},
			},
		},

		{
			[][]rune{
				[]rune{'#', '#'},
				[]rune{'#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', '#'},
				[]rune{'#', '#'},
			},
		},




		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
			},
		},

		{
			[][]rune{
				[]rune{'#', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{'#', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', '#', ' '},
			},
		},

		{
			[][]rune{
				[]rune{'#', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', '#'},
			},
		},





		{
			[][]rune{
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
				[]rune{' ', ' ', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' ', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
			},
		},

		{
			[][]rune{
				[]rune{' ', ' ', ' ', ' '},
				[]rune{' ', 'x', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{'x', ' ', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{' ', 'x', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{'x', 'x', 'x', ' '},
				[]rune{' ', ' ', 'x', ' '},
				[]rune{' ', ' ', 'x', 'x'},
			},
		},


		{
			[][]rune{
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', '#', '#', ' '},
				[]rune{' ', ' ', ' ', ' ', '#', ' ', ' ', '#'},
				[]rune{' ', ' ', '#', ' ', ' ', '#', '#', ' '},
				[]rune{' ', ' ', ' ', '#', ' ', ' ', ' ', ' '},
				[]rune{' ', '#', '#', '#', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', '#', ' ', ' ', '#', ' '},
				[]rune{' ', ' ', '#', ' ', '#', ' ', '#', ' '},
				[]rune{' ', '#', ' ', ' ', '#', ' ', '#', ' '},
				[]rune{' ', ' ', '#', '#', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{'#', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
				[]rune{'#', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				[]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			},
		},
	}


	for _,datum := range tests {

		// Figure out widith and height of the world map.
			height := len(datum.worldMap)
			width  := len(datum.worldMap[0])



		// Create the world.
			world := NewWorld(width, height)



		// Confirm widith
			actualWidth := world.Width()

			if width != actualWidth {
				t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected width of [%v].", width, height, actualWidth)
				continue
			}



		// Confirm height
			actualHeight := world.Height()

			if height != actualHeight {
				t.Errorf("When creating new world with input width,height = %vx%v, got back unexpected height of [%v].", width, height, actualHeight)
				continue
			}



		// Confirm Get() returns what was put into the world with Set()
			for h:=0; h<len(datum.worldMap); h++ {
				for w:=0; w<len(datum.worldMap[h]); w++ {

					world.Set(w,h, datum.worldMap[h][w])
				}
			}

			for h:=0; h<len(datum.worldMap); h++ {
				for w:=0; w<len(datum.worldMap[h]); w++ {

					r := world.Get(w,h)
					if r != datum.worldMap[h][w] {
						t.Errorf("When getting from world with x,y = %v,%v, expected to get back [%v] but instead got back [%v].", w, h, string(datum.worldMap[h][w]), string(r))
					}

					for xm:=1; xm<20; xm++ {
						for ym:=1; ym<20; ym++ {

							xxx := w - xm*actualWidth
							yyy := h - ym*actualHeight


							r := world.Get(xxx,yyy)
							if r != datum.worldMap[h][w] {
								t.Errorf("When getting from world with (xxx,yyy) = %v,%v, expected to get back [%v] but instead got back [%v].", xxx, yyy, string(datum.worldMap[h][w]), string(r))
							}

						}
					}

				}
			}
	}
}
