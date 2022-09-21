package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/miniscruff/igloo/mathf"
)

type Label struct {
	Transform *mathf.Transform
	font      font.Face
	text      string
	color     color.Color
	inView    bool
	options   *ebiten.DrawImageOptions
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(newText string) {
	if l.text == newText {
		return
	}

	l.text = newText
	rect := text.BoundString(l.font, l.text)
	l.Transform.SetNaturalWidth(float64(rect.Bounds().Dx()))
	l.Transform.SetNaturalHeight(float64(rect.Bounds().Dy()))
}

func (l *Label) Color() color.Color {
	return l.color
}

func (l *Label) SetColor(color color.Color) {
	if l.color == color {
		return
	}

	l.color = color
	l.options.ColorM.Reset()
	l.options.ColorM.ScaleWithColor(color)
}

func (l *Label) Font() font.Face {
	return l.font
}

func (l *Label) SetFont(newFont font.Face) {
	if l.font == newFont {
		return
	}

	l.font = newFont
	rect := text.BoundString(l.font, l.text)
	l.Transform.SetNaturalWidth(float64(rect.Bounds().Dx()))
	l.Transform.SetNaturalHeight(float64(rect.Bounds().Dy()))
}

func (l *Label) Draw(dest *ebiten.Image, camera Camera) {
	if l.Transform.IsDirty() || camera.IsDirty() {
		l.inView = camera.IsInView(l.Transform.Bounds())

		if l.inView {
			screenGeom := camera.WorldToScreen(l.Transform.GeoM())
			l.options.GeoM = screenGeom
		}
	}

	if l.inView {
		text.DrawWithOptions(dest, l.text, l.font, l.options)
	}
}

func NewLabel(
	font font.Face,
	labelText string,
	color color.Color,
	options ...mathf.TransformOption,
) *Label {
	rect := text.BoundString(font, labelText)
	width := float64(rect.Bounds().Dx())
	height := float64(rect.Bounds().Dy())

	// prepend our natural size option
	options = append([]mathf.TransformOption{
		mathf.TransformWithNaturalSize(width, height),
		mathf.TransformDrawFromBottom(),
	}, options...)
	transform := mathf.NewTransform(options...)

	label := &Label{
		font:      font,
		Transform: transform,
		text:      labelText,
		inView:    false,
		options:   &ebiten.DrawImageOptions{},
	}
	label.SetColor(color)

	return label
}
