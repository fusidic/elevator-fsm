package main

import "fmt"

type OpeningState struct {
	StateInfo
}

func (os *OpeningState) OnBegin() {
	fmt.Println("Elevator is opening")
}

func (os *OpeningState) OnEnd() {
	fmt.Println("Elevator is closing")
}

// Overwrite EnableSelfTransit()
func (os *OpeningState) EnableSelfTransit() bool {
	return true
}

// Compete state transiting road-map
func (os *OpeningState) CanTransitTo(name string) bool {
	return name == "StoppingState"
}

type StoppingState struct {
	StateInfo
}

func (ss *StoppingState) OnBegin() {
	fmt.Println("Elevator is stopped")
}

func (ss *StoppingState) OnEnd() {
	fmt.Println("Watch out, elevator is going to move")
}

func (ss *StoppingState) CanTransitTo(name string) bool {
	return name != "StoppingState"
}

type RunningState struct {
	StateInfo
}

func (rs *RunningState) OnBegin() {
	fmt.Println("Elevator is moving")
}

func (rs *RunningState) OnEnd() {
	fmt.Println("Elevator reaches target floor")
}

func (rs *RunningState) CanTransitTo(name string) bool {
	return name == "StoppingState"
}

func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED transit from %s to %s\nError: %s\n", GetStateName(sm.curr), target, err.Error())
	}
}

func main() {
	sm := NewStateManager()
	sm.OnChange = func(from, to State) {
		fmt.Printf("%s -----> %s\n", GetStateName(from), GetStateName(to))
	}

	// Instantiate states and Add them to statemgr
	sm.Add(new(OpeningState))
	sm.Add(new(StoppingState))
	sm.Add(new(RunningState))

	transitAndReport(sm, "OpeningState")
	transitAndReport(sm, "OpeningState")
	transitAndReport(sm, "StoppingState")
	transitAndReport(sm, "RunningState")
	transitAndReport(sm, "StoppingState")
	transitAndReport(sm, "OpeningState")
}