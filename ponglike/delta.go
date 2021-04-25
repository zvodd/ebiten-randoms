package main

import "time"

var(
	_DeltaTime time.Duration
	_LastTime time.Time
	Delta float64
)

func UpdateDelta(){
	_DeltaTime = time.Since(_LastTime) 
	Delta = float64(_DeltaTime / time.Millisecond)
	if Delta < 1{
		Delta = 1
	}
	_LastTime = time.Now()
}