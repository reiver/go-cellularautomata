package square



import (
	"testing"
)



func TestCanonicalXY(t *testing.T) {

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

		{3 , 5},
		{5 , 3},

		{4 , 8},
		{8 , 4},

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


		for y:=0; y<actualHeight; y++ {
			for x:=0; x<actualWidth; x++ {

				xx,yy := world.canonicalXY(x,y)

				if xx != x || yy != y {
					t.Errorf("When calculating canonicalXY() on (x,y) = (%v,%v) with width,height = %vx%v, got back unexpected (xx,yy) = (%v,%v).", x, y, datum.width, datum.height, xx, yy)
					continue
				}


				for xm:=1; xm<20; xm++ {
					for ym:=1; ym<20; ym++ {

						xxx := x + xm*actualWidth
						yyy := y + ym*actualHeight

						xx,yy := world.canonicalXY(xxx,yyy)

						if xx != x || yy != y {
							t.Errorf("When calculating canonicalXY() (xxx,yyy) = (%v,%v) with xm=%v, ym=%v, (x,y)=(%v,%v) and width,height = %vx%v, got back unexpected (xx,yy) = (%v,%v).",  xxx,yyy, xm, ym, x, y, datum.width, datum.height, xx, yy)
							continue
						}
					}
				}



				for xm:=1; xm<20; xm++ {
					for ym:=1; ym<20; ym++ {

						xxx := x - xm*actualWidth
						yyy := y - ym*actualHeight

						xx,yy := world.canonicalXY(xxx,yyy)

						if xx != x || yy != y {
							t.Errorf("When calculating canonicalXY() (xxx,yyy) = (%v,%v) with xm=-%v, ym=-%v, (x,y)=(%v,%v) and width,height = %vx%v, got back unexpected (xx,yy) = (%v,%v).",  xxx,yyy, xm, ym, x, y, datum.width, datum.height, xx, yy)
							continue
						}
					}
				}

			}
		}
	}
}



