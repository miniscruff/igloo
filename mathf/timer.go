package mathf

type TimerCompleteFunc func()

type TimerRepeatMode string

const (
	TimerNoRepeat TimerRepeatMode = "NoRepeat"
	TimerRepeats  TimerRepeatMode = "Repeats"
)

type Timer struct {
	timer        float64
	duration     float64
	isPaused     bool
	isComplete   bool
	repeat       TimerRepeatMode
	completeFunc TimerCompleteFunc
}

func NewTimer(duration float64, options ...TimerOption) *Timer {
	t := &Timer{
		timer:        0,
		duration:     duration,
		isPaused:     true,
		isComplete:   false,
		completeFunc: func() {},
	}

	for _, o := range options {
		o(t)
	}

	// set the start time here with an option to customize time.Now

	return t
}

func (t *Timer) Start() {
	t.Resume()
	t.timer = 0
}

func (t *Timer) Pause() {
	t.isPaused = true
}

func (t *Timer) Resume() {
	t.isPaused = false
}

func (t *Timer) IsPaused() bool {
	return t.isPaused
}

func (t *Timer) IsComplete() bool {
	return t.isComplete
}

func (t *Timer) Tick(gameTime *GameTime) {
	t.timer += gameTime.DeltaTime()
	if t.timer >= t.duration {
		t.completeFunc()

		if t.repeat == TimerRepeats {
			t.timer = 0
		} else {
			t.isComplete = true
		}
	}
}

type TimerOption func(timer *Timer)

func TimerOnComplete(fn TimerCompleteFunc) TimerOption {
	return func(t *Timer) {
		t.completeFunc = fn
	}
}

func TimerWithRepeat(mode TimerRepeatMode) TimerOption {
	return func(t *Timer) {
		t.repeat = mode
	}
}
