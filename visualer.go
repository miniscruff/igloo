package igloo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo/mathf"
)

type Visualer struct {
	*mathf.Transform

	Dirtier
	Drawer
	NativeSizer

	Parent   *Visualer
	Children []*Visualer
	visible  bool

	nowVisible           bool
	forcedTransformDirty bool
	forcedDirty          bool
}

func (v *Visualer) InsertChild(child *Visualer) {
	v.Children = append(v.Children, child)
	child.Parent = v
}

// ForceDirty can be called if you really need to force a rebuild of the transform
// in the next layout.
func (v *Visualer) ForceDirty() {
	v.forcedDirty = true
}

func (v *Visualer) Visible() bool {
	return v.visible
}

func (v *Visualer) SetVisible(state bool) {
	if v.visible == state {
		return
	}

	v.visible = state

	if state {
		v.nowVisible = true
	}
}

func (v *Visualer) Layout(
	offset *mathf.Transform,
	root *mathf.Transform,
) {
	if !v.visible {
		return
	}

	// if we were just turned on, set it to all our children
	if v.nowVisible {
		for _, child := range v.Children {
			child.visible = true
			child.nowVisible = true
		}
	}

	// if our transform is dirty, or our parent forced it
	// update our children as well
	if v.Transform.IsDirty() || v.forcedTransformDirty {
		v.forcedTransformDirty = true
		for _, child := range v.Children {
			child.forcedTransformDirty = true
		}
	}

	// if our own visual is dirty, force our children to be as well
	if v.Dirtier.IsDirty() || v.forcedDirty {
		for _, child := range v.Children {
			child.forcedDirty = true
		}
	}

	// our own visual changed, update our own state
	// I think a generic callback would work better here
	// something on `OnDirty() { ... }`
	if v.Dirtier.IsDirty() {
		nativeWidth, nativeHeight := v.NativeSize()
		v.Transform.SetNaturalWidth(nativeWidth)
		v.Transform.SetNaturalHeight(nativeHeight)
		v.Dirtier.Clean()
	}

	if v.nowVisible || v.forcedDirty || v.forcedTransformDirty {
		v.Transform.Build(offset)
	}

	// needs to be after we try and build
	if !root.InView(v.Transform) {
		return
	}

	// TODO: rotation and scale
	offset.Translate(v.Transform.Position())

	for _, child := range v.Children {
		child.Layout(offset, root)
	}

	// TODO: rotation and scale
	offset.Translate(v.Transform.Position().MulScalar(-1))

	v.Transform.Clean()
	v.nowVisible = false
	v.forcedDirty = false
	v.forcedTransformDirty = false
}

func (v *Visualer) Draw(dest *ebiten.Image) {
	if !v.visible {
		return
	}

	v.Drawer.Draw(dest)

	for _, child := range v.Children {
		child.Draw(dest)
	}
}
