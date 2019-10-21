package game

import "time"

type Timer struct {
	oldTime int64
	delta   int64
}

func (t *Timer) Start() {
	t.oldTime = time.Now().UnixNano()
}

func (t *Timer) Update() {
	current := time.Now().UnixNano()
	t.delta = current - t.oldTime
	t.oldTime = current
}

func (t *Timer) GetDelta() int64 {
	return t.delta
}