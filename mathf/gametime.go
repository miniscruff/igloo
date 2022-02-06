package mathf

import (
	"time"
)

type GameTime struct {
	startTime time.Time
	lastTime  time.Time
	deltaTime float64
}

func NewGameTime() *GameTime {
	return &GameTime{
		startTime: time.Now(),
		lastTime:  time.Now(),
		deltaTime: 0,
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

func (gt *GameTime) IsSlow() bool {
	// anything slower then 30fps is slow for now
	return gt.deltaTime > (1.0 / 30.0)
}

func (gt *GameTime) Tick() {
	now := time.Now()
	gt.deltaTime = now.Sub(gt.lastTime).Seconds()
	gt.lastTime = now
}
