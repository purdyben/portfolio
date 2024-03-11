package games

import (
	"context"
	"time"
)

const (
	Mill200 time.Duration = 200 * time.Millisecond
	OneSec                = 1 * time.Second
	TwoSec                = 2 * time.Second
	FiveSec               = 5 * time.Second
	TenSec                = 10 * time.Second
)

type GameTicker struct {
	tick   chan int
	delay  time.Duration
	cancel context.CancelFunc
	Ticker *time.Ticker
	paused bool
}

func NewTicker(delay time.Duration) *GameTicker {
	t := new(GameTicker)
	t.tick = make(chan int)
	t.delay = delay

	go backgroundTick(t)

	return t
}

func backgroundTick(t *GameTicker) {
	t.Ticker = time.NewTicker(t.delay)
	defer t.Ticker.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	t.cancel = cancel

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.Ticker.C:
			t.paused = false
			t.tick <- 1
		}
	}
}

func (t *GameTicker) Close() {
	t.cancel()
}

func (t *GameTicker) Tick() chan int {
	return t.tick
}

// func (t *GameTicker) SetDelay(delay time.Duration) {
// 	t.cancel()
// 	t.delay = delay
// 	go backgroundTick(t)
// }

func (t *GameTicker) Stop() {
	defer recover()
	t.Ticker.Stop()
	t.paused = true
}

func (t *GameTicker) Start() {
	defer recover()
	t.Ticker.Reset(t.delay)
}

func (t *GameTicker) Reset(d time.Duration) {
	defer recover()
	t.Ticker.Reset(d)
	// if t.paused {
	// 	t.Resume()
	// 	t.paused = false
	// }
}
