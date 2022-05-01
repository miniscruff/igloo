package graphics

import (
	"github.com/miniscruff/igloo/mathf"
)

type AnimationClip struct {
	tweens []*mathf.Tween
}

func NewAnimationClip(ticker *mathf.Ticker, tweens ...*mathf.Tween) *AnimationClip {
	for _, t := range tweens {
		ticker.Add(t)
		t.Pause()
	}

	return &AnimationClip{
		tweens: tweens,
	}
}

func (ac *AnimationClip) Start() {
	for _, t := range ac.tweens {
		t.Start()
	}
}

func (ac *AnimationClip) Pause() {
	for _, t := range ac.tweens {
		t.Pause()
	}
}

func (ac *AnimationClip) Resume() {
	for _, t := range ac.tweens {
		t.Resume()
	}
}
