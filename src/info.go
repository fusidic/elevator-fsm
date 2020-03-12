package main

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) setName(name string) {
	s.name = name
}

// Default info
func (s *StateInfo) EnableSelfTransit() bool {
	// False by default
	return false
}

func (s *StateInfo) OnBegin() {
}

func (s *StateInfo) OnEnd() {
}

func (s *StateInfo) CanTransitTo(name string) bool {
	// True by default
	return true
}

