package animation

import (
	"time"

	"fyne.io/fyne/v2"
	"go.uber.org/atomic"
)

type anim struct {
	a           *fyne.Animation
	end         time.Time
	repeatsLeft int
	reverse     bool
	start       time.Time
	total       int64

	stopped *atomic.Bool
}

func newAnim(a *fyne.Animation) *anim {
	animate := &anim{a: a, start: time.Now(), end: time.Now().Add(a.Duration)}
	animate.total = animate.end.Sub(animate.start).Nanoseconds() / 1000000 // TODO change this to Milliseconds() when we drop Go 1.12
	animate.repeatsLeft = a.RepeatCount
	return animate
}

func (a *anim) setStopped() {
	a.stopped.Store(true)
}

func (a *anim) isStopped() bool {
	return a.stopped.Load()
}
