package main

import (
	"testing"
)

func TestTrim(t *testing.T){
	s := " Mooxx New "

	expect := "Mooxx New"

	actual := trim(s)

	if actual != expect {
		t.Error("expect")
		t.Error(expect)
		t.Error("to equal")
		t.Error(actual)
	}
}

func TestLast(t *testing.T){
	s := "Mooxx New"

	expect := "New"

	actual := last(s)

	if actual != expect {
		t.Error("expect")
		t.Error(expect)
		t.Error("to equal")
		t.Error(actual)
	}
}

func TestFullName(t *testing.T){
	expect := "Mooxx New"

	actual := toFullName("Mooxx","New")

	if actual != expect {
		t.Error("expect")
		t.Error(expect)
		t.Error("to equal")
		t.Error(actual)

	}
}
