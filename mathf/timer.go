package mathf

import "time"

type TimerCompleteFunc func()

type TimerRepeatMode string

const (
	TimerNoRepeat TimerRepeatMode = "NoRepeat"
	TimerRepeats  TimerRepeatMode = "Repeats"
)

type Timer struct {
	startTime    time.Time
	duration     float64
	isPaused     bool
	isComplete   bool
	repeat       TimerRepeatMode
	completeFunc TimerCompleteFunc
}

func NewTimer(duration float64, options ...TimerOption) *Timer {
	t := &Timer{
		startTime:    time.Now(),
		duration:     duration,
		isPaused:     false,
		isComplete:   false,
		completeFunc: func() {},
	}

	for _, o := range options {
		o(t)
	}

	return t
}

func (t *Timer) Resume() {
	t.isPaused = false
}

func (t *Timer) Pause() {
	t.isPaused = true
}

func (t *Timer) IsPaused() bool {
	return t.isPaused
}

func (t *Timer) IsComplete() bool {
	return t.isComplete
}

func (t *Timer) Tick(gameTime *GameTime) {
	if time.Since(t.startTime).Seconds() >= t.duration {
		t.completeFunc()
		if t.repeat == TimerRepeats {
			t.startTime = time.Now()
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
