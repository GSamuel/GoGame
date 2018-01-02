package server

func NewEmptySlot(timeOutSeconds float64) *Slot {
	return &Slot{true, NewTimeOut(timeOutSeconds), Connection{""}}
}

type Slot struct {
	open       bool
	timeOut    TimeOut
	connection Connection
}

func (s *Slot) IsOpen() bool {
	return s.open
}

func (s *Slot) Equals(c Connection) bool {
	return s.connection == c
}

func (s *Slot) SetConnection(c Connection) {
	s.connection = c
	s.open = false
	s.ResetTimeOut()
}

func (s *Slot) ResetTimeOut() {
	s.timeOut.Reset()
}

func (s *Slot) TimedOut(timeOut float64) bool {
	return s.timeOut.Over()
}
