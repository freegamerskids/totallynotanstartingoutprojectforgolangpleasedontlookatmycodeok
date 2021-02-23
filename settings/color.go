package settings

var Colors = InitColors()

func InitColors() *colors {
	return &colors{
		RED:   0,
		GREEN: 0,
		BLUE:  0,
		ALPHA: 1,
	}
}

type colors struct {
	RED   uint8
	GREEN uint8
	BLUE  uint8
	ALPHA float32
}

func (clr *colors) setColor(r uint8, g uint8, b uint8, a float32) {
	clr.RED = r
	clr.GREEN = g
	clr.BLUE = b
	clr.ALPHA = a
}
