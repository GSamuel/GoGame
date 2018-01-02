package server

import (
	"fmt"
)

type Server struct {
	*ServerSettings
	Slots []*Slot
}

func (s *Server) FindSlot(c Connection) (int, error) {
	for i := 0; i < s.MaxConnections; i++ {
		if s.Slots[i].Equals(c) && !s.Slots[i].IsOpen() {
			return i, nil
		}

	}
	return -1, fmt.Errorf("Connection does not exist")
}

func (s *Server) FindOpenSlot() (int, error) {
	for i := 0; i < s.MaxConnections; i++ {
		if s.Slots[i].IsOpen() {
			return i, nil
		}
	}

	return -1, fmt.Errorf("No free slots available")
}

func New(settings ServerSettings) *Server {
	var slots = make([]*Slot, settings.MaxConnections, settings.MaxConnections)

	for i := 0; i < len(slots); i++ {
		slots[i] = NewEmptySlot(settings.Timeout)
	}

	return &Server{&settings, slots}
}

func (s *Server) ResolveConnection(c Connection) error {
	i, err := s.FindSlot(c)

	if err == nil {
		fmt.Println("Packet received from ", c.Address(), " updating timestamp")
		s.Slots[i].ResetTimeOut()

		return nil
	}

	fmt.Println("New connection ", c.Address())

	i, err = s.FindOpenSlot()

	if err != nil {
		fmt.Println("No available slots")
		return err
	}

	fmt.Println("New connection now in slot ", i)
	s.Slots[i].SetConnection(c)

	return nil

	//does connection already exist?
	//y: update with current time, return;
	//find open slot:
	//put connection in open slot.
	//set current time.
}
