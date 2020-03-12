package main

import "reflect"

type State interface {
	Name() string
	EnableSelfTransit() bool
	OnBegin()
	OnEnd()
	CanTransitTo(name string) bool
}

func GetStateName(s State) string {
	if s == nil {
		return "none"
	}
	// Use reflect to get the name of state (not the var name)
	return reflect.TypeOf(s).Elem().Name()
}