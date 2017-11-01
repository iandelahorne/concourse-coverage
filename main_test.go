package main

import "testing"

func TestDoSomething(t *testing.T) {
	arg := 2
	expected := 34
	actual := doSomething(arg)
	if actual != expected {
		t.Fail()
	}
}

func TestOtherThing(t *testing.T) {
	arg := 2 
	expected := 15
	actual := doOtherThing(arg)

	if actual != expected {
		t.Fail()
	}	
}