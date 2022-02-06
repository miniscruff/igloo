package mathf

type TweenOption func(t *Tween)

func TweenWithEase(easeFunc EaseFunc) TweenOption {
	return func(t *Tween) {
		t.easeFunc = easeFunc
	}
}

func TweenUpdatePointer(value *float64) TweenOption {
	return func(t *Tween) {
		t.valueFunc = func(newValue float64) {
			*value = newValue
		}
	}
}

func TweenUpdateFunc(fn ValueFunc) TweenOption {
	return func(t *Tween) {
		t.valueFunc = fn
	}
}

func TweenWithRepeat(mode RepeatMode) TweenOption {
	return func(t *Tween) {
		t.repeat = mode
	}
}

type ValueFunc func(value float64)

type RepeatMode string

const (
	// Does not repeat, will trigger complete and be removed from ticker
	NoRepeat RepeatMode = "NoRepeat"
	// Repeats the tween from the start over and over again
	RepeatLoop RepeatMode = "Loop"
	// Bounces the tween going from 0>1>0, reversing the ease to get a good bounce
	RepeatBounce RepeatMode = "Bounce"
	// Pauses the tween at the end ready to start again
	RepeatPause RepeatMode = "Pause"
)

type Tween struct {
	duration   float64
	percent    float64
	isPaused   bool
	isComplete bool
	easeFunc   EaseFunc
	valueFunc  ValueFunc
	repeat     RepeatMode
	isBouncing bool
}

func NewTween(duration float64, options ...TweenOption) *Tween {
	t := &Tween{
		duration:   duration,
		percent:    0,
		isPaused:   false,
		easeFunc:   EaseLinear,
		valueFunc:  func(value float64) {},
		repeat:     NoRepeat,
		isBouncing: false,
	}

	for _, o := range options {
		o(t)
	}

	return t
}

func (t *Tween) Start() {
	t.isPaused = false
}

func (t *Tween) Pause() {
	t.isPaused = true
}

func (t *Tween) IsPaused() bool {
	return t.isPaused
}

func (t *Tween) IsComplete() bool {
	return t.isComplete
}

func (t *Tween) Tick(gameTime *GameTime) {
	t.percent += gameTime.DeltaTime() / t.duration
	if t.percent >= 1 {
		t.percent = 1

		// handle completion based on repeat mode
		switch t.repeat {
		case NoRepeat:
			t.isComplete = true
		case RepeatLoop:
			t.percent = 0
		case RepeatBounce:
			t.isBouncing = !t.isBouncing
			t.percent = 0
		case RepeatPause:
			t.isPaused = true
		}
	}

	easedPercent := t.easeFunc(t.percent)
	if t.isBouncing {
		easedPercent = t.easeFunc(1 - t.percent)
	}

	t.valueFunc(easedPercent)
}

func TweenVec2Func(start, end Vec2, fn func(Vec2)) ValueFunc {
	return func(value float64) {
		fn(Vec2Lerp(start, end, value))
	}
}

func TweenVec2Pointer(start, end Vec2, ptr *Vec2) ValueFunc {
	return func(value float64) {
		*ptr = Vec2Lerp(start, end, value)
	}
}
