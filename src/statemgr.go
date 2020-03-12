package main

import (
	"errors"
)

type StateManager struct {
	stateByName map[string]State
	// Callback while state gets changed
	OnChange func(from, to State)
	curr State
}

func (sm *StateManager) Get(name string) State {
	// In case get panic when 'name' do not exist
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

func (sm *StateManager) Add(s State) {
	name := GetStateName(s)
	// State does not have any interface of setName()
	// Though StateInfo gets one
	s.(interface{
		// Actually this method is from StateInfo
		setName(name string)
	}).setName(name)

	// Check name duplication
	if sm.Get(name) != nil {
		panic("Duplicate state:" + name)
	}

	sm.stateByName[name] = s
}

// Init StateManager
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

// Define errors
var ErrStateNotFound = errors.New("state not found")
var ErrForbidSelfStateTransit = errors.New("forbid self transit")
var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) GetCurrState() State {
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	if sm.curr.Name() == name && !sm.curr.EnableSelfTransit() {
		return false
	}
	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {
	next := sm.stateByName[name]
	pre :=sm.curr
	if next == nil {
		return ErrStateNotFound
	}
	if pre != nil {
		// Maybe it's more secure to compare with string
		if GetStateName(pre) == name && pre.EnableSelfTransit() {
			goto transitHere
		}
		if !pre.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		pre.OnEnd()
	}

transitHere:
	sm.curr = next
	sm.curr.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}