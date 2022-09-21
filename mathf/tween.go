package mathf

type TweenStatusChangedFunc func(tween *Tween)
type TweenValueFunc func(value float64)

type TweenRepeatMode string

const (
	// Does not repeat, will trigger complete and be removed from ticker
	TweenNoRepeat TweenRepeatMode = "NoRepeat"
	// Repeats the tween from the start over and over again
	TweenRepeatLoop TweenRepeatMode = "Loop"
	// Bounces the tween going from 0>1>0 in a loop, reversing the ease to get a good bounce
	TweenRepeatBounceLoop TweenRepeatMode = "RepeatBounce"
	// Bounces the tween going from 0>1>0, reversing the ease to get a good bounce
	TweenRepeatBounce TweenRepeatMode = "Bounce"
	// Pauses the tween at the end ready to start again
	TweenRepeatPause TweenRepeatMode = "Pause"
)

type Tween struct {
	duration   float64
	percent    float64
	isPaused   bool
	isComplete bool
	easeFunc   EaseFunc
	valueFunc  TweenValueFunc
	repeat     TweenRepeatMode
	started    TweenStatusChangedFunc
	completed  TweenStatusChangedFunc
	isBouncing bool
	firstStart bool
}

func NewTween(duration float64, options ...TweenOption) *Tween {
	t := &Tween{
		duration:   duration,
		percent:    0,
		isPaused:   true,
		easeFunc:   EaseLinear,
		valueFunc:  func(value float64) {},
		started:    func(t *Tween) {},
		completed:  func(t *Tween) {},
		repeat:     TweenNoRepeat,
		isBouncing: false,
		firstStart: true,
	}

	for _, o := range options {
		o(t)
	}

	return t
}

func (t *Tween) Start() {
	t.Resume()
	t.started(t)
	t.firstStart = false
}

// Bounce will start the tween bouncing back if the
// tween was not yet complete
func (t *Tween) Bounce() {
	if !t.isPaused && !t.firstStart {
		t.isBouncing = !t.isBouncing
	}
	t.Start()
}

func (t *Tween) Pause() {
	t.isPaused = true
}

func (t *Tween) Resume() {
	t.isPaused = false
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
		case TweenNoRepeat:
			t.isComplete = true
		case TweenRepeatLoop:
			t.percent = 0
		case TweenRepeatBounceLoop:
			t.isBouncing = !t.isBouncing
			t.percent = 0
		case TweenRepeatBounce:
			t.isBouncing = !t.isBouncing
			t.percent = 0
			t.isPaused = true
		case TweenRepeatPause:
			t.isPaused = true
		}

		t.completed(t)
	}

	easedPercent := t.easeFunc(t.percent)
	if t.isBouncing {
		easedPercent = t.easeFunc(1 - t.percent)
	}

	t.valueFunc(easedPercent)
}

func TweenVec2Func(start, end Vec2, fn func(Vec2)) TweenValueFunc {
	return func(value float64) {
		fn(Vec2Lerp(start, end, value))
	}
}

func TweenVec2Pointer(start, end Vec2, ptr *Vec2) TweenValueFunc {
	return func(value float64) {
		*ptr = Vec2Lerp(start, end, value)
	}
}

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

func TweenUpdateFunc(fn TweenValueFunc) TweenOption {
	return func(t *Tween) {
		t.valueFunc = fn
	}
}

func TweenWithRepeat(mode TweenRepeatMode) TweenOption {
	return func(t *Tween) {
		t.repeat = mode
	}
}

func TweenOnStart(fn TweenStatusChangedFunc) TweenOption {
	return func(t *Tween) {
		t.started = fn
	}
}

func TweenOnComplete(fn TweenStatusChangedFunc) TweenOption {
	return func(t *Tween) {
		t.completed = fn
	}
}

func TweenPlay() TweenOption {
	return func(t *Tween) {
		t.isPaused = false
	}
}
