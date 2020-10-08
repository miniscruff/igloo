package igloo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/miniscruff/igloo"
)

var _ = Describe("Transform", func() {
	var transform *igloo.Transform

	BeforeEach(func() {
		transform = igloo.NewTransform(100, 200, 3.14)
		transform.Clean() // default to a clean transform
	})

	It("New creates dirty transform", func() {
		dirtyTransform := igloo.NewTransform(100, 200, 3.14)
		Expect(dirtyTransform.IsDirty()).To(Equal(true))
		Expect(dirtyTransform.X()).To(Equal(100.0))
		Expect(dirtyTransform.Y()).To(Equal(200.0))
		Expect(dirtyTransform.Rotation()).To(Equal(3.14))
	})

	It("can be cleaned", func() {
		// done in before each
		// transform.Clean()
		Expect(transform.IsDirty()).To(Equal(false))
	})

	It("can set x", func() {
		transform.SetX(12.5)
		Expect(transform.IsDirty()).To(Equal(true))
		Expect(transform.X()).To(Equal(12.5))
	})

	It("can set y", func() {
		transform.SetY(30.3)
		Expect(transform.IsDirty()).To(Equal(true))
		Expect(transform.Y()).To(Equal(30.3))
	})

	It("can set rotation", func() {
		transform.SetRotation(0.707)
		Expect(transform.IsDirty()).To(Equal(true))
		Expect(transform.Rotation()).To(Equal(0.707))
	})

	It("can set x and y", func() {
		transform.SetPosition(24.4, 32.3)
		Expect(transform.IsDirty()).To(Equal(true))
		x, y := transform.GetPosition()
		Expect(x).To(Equal(24.4))
		Expect(y).To(Equal(32.3))
	})

	It("can move a transform", func() {
		transform.Translate(50, 50)
		Expect(transform.IsDirty()).To(Equal(true))
		Expect(transform.X()).To(Equal(150.0))
		Expect(transform.Y()).To(Equal(250.0))
	})
})
