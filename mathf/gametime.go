package mathf

import (
	"time"
)

type GameTime struct {
	startTime time.Time
	deltaTime float64
}

func NewGameTime() *GameTime {
	return &GameTime{
		startTime: time.Now(),
		deltaTime: 1.0/60.0,
	}
}

func (gt *GameTime) TotalDuration() time.Duration {
	return time.Since(gt.startTime)
}

func (gt *GameTime) TotalSeconds() float64 {
	return gt.TotalDuration().Seconds()
}

func (gt *GameTime) DeltaTime() float64 {
	return gt.deltaTime
}
