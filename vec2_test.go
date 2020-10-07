package igloo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"

	"github.com/miniscruff/igloo"
)

var _ = Describe("Vec2", func() {
	It("has string method", func() {
		a := &igloo.Vec2{X: 10, Y: 20}
		Expect(a.String()).To(Equal("Vec2(10, 20)"))
	})

	It("can add two vectors", func() {
		a := igloo.Vec2{X: 1, Y: 2}
		b := igloo.Vec2{X: 4, Y: 6}
		Expect(a.Add(b)).To(Equal(igloo.Vec2{X: 5, Y: 8}))
	})

	It("can multiply a vector", func() {
		a := igloo.NewVec2(10, 20)
		Expect(a.MulScalar(3)).To(Equal(igloo.NewVec2(30, 60)))
	})

	It("can subtract two vectors", func() {
		a := igloo.NewVec2(12, 6)
		b := igloo.NewVec2(4, 3)
		Expect(a.Sub(b)).To(Equal(igloo.NewVec2(8, 3)))
	})

	It("can subtract a scalar", func() {
		a := igloo.NewVec2(12, 6)
		Expect(a.SubScalar(5)).To(Equal(igloo.NewVec2(7, 1)))
	})

	It("can unit zero", func() {
		Expect(igloo.NewVec2(0, 0).Unit()).To(Equal(igloo.Vec2Zero))
	})

	It("can unit a vector", func() {
		a := igloo.NewVec2(4, 7)
		unit := igloo.Vec2{
			X: 0.49613893835683387,
			Y: 0.8682431421244593,
		}
		Expect(a.Unit()).To(Equal(unit))
		Expect(unit.Mag()).To(BeNumerically("~", 1.0))
		Expect(unit.SqrMag()).To(BeNumerically("~", 1.0))
	})

	It("can get magnitude", func() {
		a := igloo.Vec2{X: 5, Y: 12}
		Expect(a.Mag()).To(Equal(13.0))
	})

	It("can get squared magnitude", func() {
		a := igloo.Vec2{X: 3, Y: 4}
		Expect(a.SqrMag()).To(Equal(25.0))
	})

	It("can get distance between two vectors", func() {
		a := igloo.NewVec2(7, 7)
		b := igloo.NewVec2(4, 3)
		Expect(a.Dist(b)).To(Equal(5.0))
	})

	It("can get squared distance between two vectors", func() {
		a := igloo.NewVec2(7, 7)
		b := igloo.NewVec2(4, 3)
		Expect(a.SqrDist(b)).To(Equal(25.0))
	})

	It("can get X and Y separately", func() {
		a := igloo.Vec2{X: 1, Y: 2}
		x, y := a.XY()
		Expect(x).To(Equal(1.0))
		Expect(y).To(Equal(2.0))
	})

	It("can get the angle", func() {
		a := igloo.Vec2{X: 4, Y: 5}
		Expect(a.Angle()).To(BeNumerically("~", 0.8960553845713439))
	})

	It("can get vector from angle", func() {
		a := igloo.Vec2FromAngle(math.Pi / 4)
		Expect(a.X).To(BeNumerically("~", 0.7071067812))
		Expect(a.Y).To(BeNumerically("~", 0.7071067812))
	})

	It("can get the normal", func() {
		a := igloo.Vec2{X: 12, Y: 5}
		Expect(a.Normal()).To(Equal(igloo.Vec2{X: -5, Y: 12}))
	})

	It("can get the dot product", func() {
		a := igloo.Vec2{X: 2, Y: 3}
		b := igloo.Vec2{X: 4, Y: 5}
		Expect(a.Dot(b)).To(Equal(23.0))
	})

	It("can get the cross product", func() {
		a := igloo.Vec2{X: 8, Y: 4}
		b := igloo.Vec2{X: 3, Y: 3}
		Expect(a.Cross(b)).To(Equal(12.0))
	})

	It("can map a func to values", func() {
		a := igloo.Vec2{X: 10, Y: 12}
		addTen := func(v float64) float64 {
			return v + 10
		}
		b := igloo.Vec2{X: 20, Y: 22}
		Expect(a.Map(addTen)).To(Equal(b))
	})
})
