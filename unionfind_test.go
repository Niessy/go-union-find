package gofind

import (
	"testing"
)

func intCompare(a, b interface{}) bool {
	return a.(int) < b.(int)
}

type T struct {
	u, v, x int
}

func TestMakeSetSpace(t *testing.T) {
	ss := MakeSetSpace(intCompare)
	ss.AddorUpdateSet("first", 10)
	ss.AddorUpdateSet("second", 5)
	ss.AddorUpdateSet("third", 50)
	ss.AddorUpdateSet("fourth", 22)
	if ss.Count() != 4 {
		t.Errorf("SetSpace count should be 4 instead it's %d.\n", ss.Count())
	}
}

func TestFindAndUnion(t *testing.T) {
	ss := MakeSetSpace(intCompare)
	ss.AddorUpdateSet("first", 10)
	ss.AddorUpdateSet("second", 5)
	ss.AddorUpdateSet("third", 50)
	ss.AddorUpdateSet("fourth", 22)
	ss.UnionbyCompare("first", "third")
	if ss.Count() != 3 {
		t.Errorf("SetSpace count should be 3 instead it's %d.\n", ss.Count())
	}
	setMap := ss.GetSetMap()
	first := setMap["first"]
	second := setMap["second"]
	third := setMap["third"]
	if first.leader != third.leader {
		t.Errorf("first should have same leader as third, first leader %v, third leader %v.\n",
			first.leader, third.leader)
	}
	ss.UnionbyCompare("second", "third")
	if second.leader != third.leader {
		t.Errorf("second should have same leader as third, second leader %v, third leader %v.\n",
			second.leader, third.leader)
	}
}
