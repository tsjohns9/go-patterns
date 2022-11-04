package adapter

import "math"

type RoundHole struct {
	radius int
}

type Radius interface {
	getRadius() int
}

func (rh *RoundHole) getRadius() int {
	return rh.radius
}

func (rh *RoundHole) fits(peg Radius) bool {
	return rh.getRadius() >= peg.getRadius()
}

type RoundPeg struct {
	radius int
}

func (rp *RoundPeg) getRadius() int {
	return rp.radius
}

type SquarePeg struct {
	width int
}

func (sp *SquarePeg) getWidth() int {
	return sp.width
}

type SquarePegAdapter struct {
	peg *SquarePeg
}

func (spa *SquarePegAdapter) getRadius() int {
	return (spa.peg.width) * int(math.Sqrt(2)/2)
}

func NewSquarePegAdapter(peg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{peg: peg}
}

func main() {
	rhole := &RoundHole{radius: 5}
	rpeg := &RoundPeg{radius: 5}
	rhole.fits(rpeg)

	smallSqpeg := &SquarePeg{width: 5}
	largeSqpeg := &SquarePeg{width: 10}

	// wrap the square peg so that it has the getRadius method
	smallSqpegAdapter := NewSquarePegAdapter(smallSqpeg)
	largeSqpegAdapter := NewSquarePegAdapter(largeSqpeg)

	rhole.fits(smallSqpegAdapter)
	rhole.fits(largeSqpegAdapter)
}
