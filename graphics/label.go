package graphics
/*

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/miniscruff/igloo/mathf"
)

// probably want a text render mode like scale, word wrap or something
// otherwise the size is just based on the text input

type Label struct {
	Transform *mathf.Transform
	Visible   bool

	font        font.Face
	text        string
	color       color.Color
	lastVisible bool
	inView      bool
	options     *ebiten.DrawImageOptions
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
	l.Transform.ResetScale()
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
	l.Transform.ResetScale()
}

func (l *Label) Draw(dest *ebiten.Image, camera Camera) {
	turnedOn := l.Visible && !l.lastVisible
	l.lastVisible = l.Visible

	if !l.Visible {
		return
	}

	if turnedOn || l.Transform.IsDirty() || camera.IsDirty() {
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
	lineHeight := float64(text.BoundString(font, "A").Dy())

	// prepend our natural size option
	options = append([]mathf.TransformOption{
		mathf.TransformWithNaturalSize(width, height),
		mathf.TransformWithFixedOffset(lineHeight),
	}, options...)
	transform := mathf.NewTransform(options...)

	label := &Label{
		Transform:   transform,
		Visible:     true,
		font:        font,
		text:        labelText,
		inView:      false,
		lastVisible: true,
		options:     &ebiten.DrawImageOptions{},
	}
	label.SetColor(color)

	return label
}
*/
