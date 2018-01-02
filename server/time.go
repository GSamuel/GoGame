package server

import (
	"time"
)

type TimeOut interface {
	Reset()
	Over() bool
}

type RealTimeOut struct {
	timeStamp        time.Time
	timoutOutSeconds float64
}

func (r *RealTimeOut) Reset() {
	r.timeStamp = time.Now()
}

func (r *RealTimeOut) Over() bool {
	return time.Since(r.timeStamp).Seconds() > r.timoutOutSeconds
}

func NewTimeOut(timeOutSeconds float64) TimeOut {
	return &RealTimeOut{time.Now(), timeOutSeconds}
}

type FakeTimeOut struct {
}
