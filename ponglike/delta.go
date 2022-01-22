package main

import "time"

type DeltaTimer struct {
	_DeltaTime time.Duration
	_LastTime  time.Time
	Delta      float64
}

func (d *DeltaTimer) Update() float64 {
	d._DeltaTime = time.Since(d._LastTime)
	d.Delta = float64(d._DeltaTime / time.Millisecond)
	if d.Delta < 1 {
		d.Delta = 1
	}
	d._LastTime = time.Now()
	return d.Delta
}
