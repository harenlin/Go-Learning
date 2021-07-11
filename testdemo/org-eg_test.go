package test_eg

import "testing"

func TestAdd(t *testing.T) {
	n := Add(8, 9)
	if n == 17 {
		t.Log("Add Success")
	} else {
		t.Error("Add Fail")
	}
}
func TestFound(t *testing.T) {
	if isSuccess := Found("Haren", []string{"Haren", "Jimmy", "Michael", "Andy"}); isSuccess {
		t.Log("Found success")
	} else {
		t.Error("Found fail")
	}
}
func TestNotFound(t *testing.T) {
	if isSuccess := Found("Steven", []string{"Haren", "Jimmy", "Michael", "Andy"}); isSuccess {
		t.Error("Found fail")
	} else {
		t.Log("Found success")
	}
}