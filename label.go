package igloo

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type LabelOptions struct {
	Font      font.Face
	Transform *Transform
	Color     color.Color
	Text      string
	Anchor    Vec2f
}

type Label struct {
	Transform *Transform
	Color     color.Color
	// dirty flagging vars
	font     font.Face
	text     string
	anchor   Vec2f
	position Vec2i
	// cache our width and height from bounds
	width  float64
	height float64
	// if we change text we need to update bounds and anchoring
	textDirty bool
	// if we change position or anchoring we dont need to recalculate our bounds
	locationDirty bool
	inView        bool
}

// IsDirty returns whether or not our object is dirty
func (l *Label) IsDirty() bool {
	return l.textDirty || l.locationDirty
}

// Clean the object back to a fresh state
func (l *Label) Clean() {
	l.textDirty = false
	l.locationDirty = false
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(newText string) {
	if l.text != newText {
		l.text = newText
		l.textDirty = true
	}
}

func (l *Label) Font() font.Face {
	return l.font
}

func (l *Label) SetFont(newFont font.Face) {
	if l.font != newFont {
		l.font = newFont
		l.textDirty = true
	}
}

func (l *Label) Width() float64 {
	return l.width
}

// Height of the rendered text, set from font size
func (l *Label) Height() float64 {
	return l.height
}

// X is the current screen position X, to move the label use Transform
func (l *Label) X() int {
	return l.position.X
}

// Y is the current screen position Y, to move the label use Transform
func (l *Label) Y() int {
	return l.position.Y
}

func (l *Label) Position() Vec2i {
	return l.position
}

func (l *Label) Anchor() Vec2f {
	return l.anchor
}

func (l *Label) SetAnchor(newAnchor Vec2f) {
	if l.anchor != newAnchor {
		l.anchor = newAnchor
		l.locationDirty = true
	}
}

func (l *Label) cacheText() {
	if !l.textDirty {
		return
	}

	// cache our width and height for anchor/position checks
	rect := text.BoundString(l.font, l.text)
	l.width = float64(rect.Bounds().Dx())
	l.height = float64(rect.Bounds().Dy())
	l.textDirty = false
}

func (l *Label) cachePosition() {
	if !l.locationDirty && !l.Transform.IsDirty() {
		return
	}

	l.position = Vec2i{
		X: int(l.Transform.X() - l.width*l.anchor.X),
		// text is drawn from the bottom so we have to 1-anchor
		Y: int(l.Transform.Y() + l.height*(1-l.anchor.Y)),
	}
	l.locationDirty = false
}

func (l *Label) Draw(dest *ebiten.Image) {
	l.cacheText()
	l.cachePosition()
	text.Draw(dest, l.text, l.font, l.X(), l.Y(), l.Color)
}

func NewLabel(options LabelOptions) (*Label, error) {
	l := &Label{
		Transform:     options.Transform,
		Color:         options.Color,
		font:          options.Font,
		text:          options.Text,
		anchor:        options.Anchor,
		width:         0,
		height:        0,
		textDirty:     true,
		locationDirty: true,
		inView:        false,
	}
	l.cacheText()
	l.cachePosition()

	return l, nil
}
