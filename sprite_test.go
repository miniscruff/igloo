package igloo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/miniscruff/igloo"
)

func emptyImage() *ebiten.Image {
	loadedImage := ebiten.NewImage(10, 10)
	loadedImage.Fill(color.RGBA{255, 255, 255, 255})

	return loadedImage
}

var blank = emptyImage()

type mockCamera struct {
	isDirty       bool
	worldToScreen func(geom ebiten.GeoM) ebiten.GeoM
	screenToWorld func(x, y float64) (float64, float64)
	isInView      func(x, y, width, height float64) bool
}

func (m *mockCamera) IsDirty() bool {
	return m.isDirty
}

func (m *mockCamera) Clean() {
	m.isDirty = false
}

func (m *mockCamera) WorldToScreen(geom ebiten.GeoM) ebiten.GeoM {
	return m.worldToScreen(geom)
}

func (m *mockCamera) ScreenToWorld(x, y float64) (float64, float64) {
	return m.screenToWorld(x, y)
}

func (m *mockCamera) IsInView(x, y, width, height float64) bool {
	return m.isInView(x, y, width, height)
}

type mockCanvas struct {
	drawImage func(src *ebiten.Image, op *ebiten.DrawImageOptions)
}

func (m *mockCanvas) DrawImage(src *ebiten.Image, op *ebiten.DrawImageOptions) {
	m.drawImage(src, op)
}

var _ = Describe("Sprite", func() {
	var transform *igloo.Transform

	BeforeEach(func() {
		transform = igloo.NewTransform(0, 0, 0)
	})

	It("New creates dirty sprite with image size", func() {
		sprite := igloo.NewSprite(blank, transform)
		Expect(sprite.IsDirty()).To(BeTrue())
		Expect(sprite.Width()).To(Equal(10.0))
		Expect(sprite.Height()).To(Equal(10.0))
	})

	It("New anchor sets the anchor", func() {
		anchor := igloo.Vec2{X: 0.25, Y: 0.25}
		sprite := igloo.NewSpriteAnchor(blank, transform, anchor)
		Expect(sprite.Anchor()).To(Equal(anchor))
	})

	It("New size sets the size", func() {
		sprite := igloo.NewSpriteSize(blank, transform, 50, 60)
		Expect(sprite.Width()).To(Equal(50.0))
		Expect(sprite.Height()).To(Equal(60.0))
	})

	It("New anchor size sets anchor and size", func() {
		anchor := igloo.Vec2{X: 0.5, Y: 0.35}
		sprite := igloo.NewSpriteAnchorSize(blank, transform, anchor, 50, 60)
		Expect(sprite.Anchor()).To(Equal(anchor))
		Expect(sprite.Width()).To(Equal(50.0))
		Expect(sprite.Height()).To(Equal(60.0))
	})

	It("can be cleaned", func() {
		dirtySprite := igloo.NewSprite(blank, transform)
		dirtySprite.Clean()
		Expect(dirtySprite.IsDirty()).To(BeFalse())
	})

	It("can change width", func() {
		sprite := igloo.NewSprite(blank, transform)
		sprite.Clean()
		sprite.SetWidth(20)
		Expect(sprite.IsDirty()).To(BeTrue())
		Expect(sprite.Width()).To(Equal(20.0))
	})

	It("can change height", func() {
		sprite := igloo.NewSprite(blank, transform)
		sprite.Clean()
		sprite.SetHeight(30)
		Expect(sprite.IsDirty()).To(BeTrue())
		Expect(sprite.Height()).To(Equal(30.0))
	})

	It("can change anchor", func() {
		anchor := igloo.Vec2{X: 0.5, Y: 0.0}
		sprite := igloo.NewSprite(blank, transform)
		sprite.Clean()
		sprite.SetAnchor(anchor)
		Expect(sprite.IsDirty()).To(BeTrue())
		Expect(sprite.Anchor()).To(Equal(anchor))
	})

	It("will draw on canvas updating all dirty parts", func() {
		calledMethods := 0
		identityOptions := &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		}
		sprite := igloo.NewSprite(blank, transform)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(src).To(Equal(blank))
				Expect(op).To(Equal(identityOptions))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				calledMethods++
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				calledMethods++
				Expect(x).To(Equal(0.0))
				Expect(y).To(Equal(0.0))
				Expect(width).To(Equal(10.0))
				Expect(height).To(Equal(10.0))
				return true
			},
		}
		Expect(sprite.IsDirty()).To(BeTrue())

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(3))
		Expect(sprite.IsDirty()).To(BeFalse())
	})

	It("will not draw or get screen position if not in view", func() {
		calledMethods := 0
		sprite := igloo.NewSprite(blank, transform)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				panic("should not draw when not in view")
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				panic("should not get screen position if not in view")
			},
			isInView: func(x, y, width, height float64) bool {
				calledMethods++
				return false
			},
		}
		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})

	It("will not check view or get screen position if nothing is dirty", func() {
		calledMethods := 0
		transform.Clean()

		sprite := igloo.NewSprite(blank, transform)
		sprite.Clean()
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				panic("should not draw when not in view")
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				panic("should not get screen position when clean")
			},
			isInView: func(x, y, width, height float64) bool {
				panic("should not check in view when clean")
			},
			isDirty: false,
		}

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(0))
	})

	It("will draw at old position if nothing is dirty", func() {
		calledMethods := 0
		identityOptions := &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		}
		sprite := igloo.NewSprite(blank, transform)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(op).To(Equal(identityOptions))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				return true
			},
		}

		// draw the object once so that the inView flag is set
		sprite.Draw(canvas, camera)

		calledMethods = 0
		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})

	It("will properly translate", func() {
		calledMethods := 0
		geo := ebiten.GeoM{}
		geo.Translate(10, 20)
		transform.SetPosition(10, 20)
		sprite := igloo.NewSprite(blank, transform)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(op.GeoM).To(Equal(geo))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				return true
			},
		}

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})

	It("will properly translate when using anchor", func() {
		calledMethods := 0
		geo := ebiten.GeoM{}
		geo.Translate(10, 10)
		transform.SetPosition(15, 15)
		sprite := igloo.NewSpriteAnchor(blank, transform, igloo.Vec2{X: 0.5, Y: 0.5})
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(op.GeoM).To(Equal(geo))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				return true
			},
		}

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})

	It("will properly scale", func() {
		calledMethods := 0
		geo := ebiten.GeoM{}
		geo.Scale(2, 2)
		geo.Translate(20, 20)

		transform.SetPosition(20, 20)
		sprite := igloo.NewSpriteSize(blank, transform, 20, 20)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(op.GeoM).To(Equal(geo))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				return true
			},
		}

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})

	It("will properly rotate", func() {
		calledMethods := 0
		geo := ebiten.GeoM{}
		geo.Rotate(3.14)

		transform.SetRotation(3.14)

		sprite := igloo.NewSprite(blank, transform)
		canvas := &mockCanvas{
			drawImage: func(src *ebiten.Image, op *ebiten.DrawImageOptions) {
				calledMethods++
				Expect(op.GeoM).To(Equal(geo))
			},
		}
		camera := &mockCamera{
			worldToScreen: func(geom ebiten.GeoM) ebiten.GeoM {
				return geom
			},
			isInView: func(x, y, width, height float64) bool {
				return true
			},
		}

		sprite.Draw(canvas, camera)
		Expect(calledMethods).To(Equal(1))
	})
})
