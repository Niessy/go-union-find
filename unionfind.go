// This is a basic implementation
// of the Union Find Data Structure.

package gofind

// Set structure
type Set struct {
	leader *Set        // pointer to leader Cluster
	rank   int         // rank of the cluster
	data   interface{} // whatever you're trying to compute with 
}

// Create a Set given data, this can be any type.
func makeSet(d interface{}) *Set {
	ns := new(Set)
	ns.rank = 0
	ns.leader = ns
	ns.data = d
	return ns
}

// A slice of Set structs
type Sets map[interface{}]*Set

// Returns length of Sets
func (c Sets) Len() int {
	return len(c)
}

type compareFunc func(a, b interface{}) bool

// Contains Sets, count (number of Clusters) and
// the comparison function for the data in the Clusters.
type SetSpace struct {
	setMap  Sets
	count   int         // count of current seperated Clusters
	compare compareFunc // function to compare Data
}

// Create a Set Space, this will contains the comparison function,
// the number of sets(that haven't been unioned!) and the sets.
func MakeSetSpace(cf compareFunc) *SetSpace {
	return &SetSpace{
		setMap:  make(Sets),
		count:   0,
		compare: cf,
	}
}

// If there is no set present with the name then it will create
// a new set and add it to the SetSpace, else it will override the
// previous set assigned to said name.
func (ss *SetSpace) AddorUpdateSet(name, data interface{}) {
	x0 := *ss
	ns := makeSet(data)
	x0.setMap[name] = ns
	x0.count++
	*ss = x0
	return
}

//Returns the number of Clusters
func (ss SetSpace) Count() int {
	return ss.count
}

//Returns the Sets map contained in the SetSpace
func (ss SetSpace) GetSetMap() Sets {
	return ss.setMap
}

// Find the leader of a Cluster
// Uses path compression so the Operation runs
// in amortized O(alpha(n)) where alpha is 
// the inverse Ackermann function
func (ss *SetSpace) find(s *Set) *Set {
	if s.leader != s {
		s.leader = ss.find(s.leader)
	}
	return s.leader
}

// Unions two Set structs, the Set with the highest rank
// is chosen for leadership.
func (ss *SetSpace) Union(a, b interface{}) {
	// find the leaders of s1 and s2 Sets
	s1 := ss.setMap[a]
	s2 := ss.setMap[b]
	c1 := ss
	x1 := c1.find(s1)
	x2 := c1.find(s2)
	// Already in the same Set
	if x2 == x1 {
		return
	}
	a1 := *x1
	a2 := *x2

	// Merge based on rank of Set
	if a1.rank < a2.rank {
		a1.leader = a2.leader
	} else if a1.rank > a2.rank {
		a2.leader = a1.leader
	} else {
		a2.leader = a1.leader
		a1.rank = a2.rank + 1
	}
	*x1 = a1
	*x2 = a2
	c1.count -= 1
	return
}

// Takes two Set structs, works like Union except that
// the predefined compareFunc is used to determine if 
// the two Set should Union.
func (ss *SetSpace) UnionbyCompare(a, b interface{}) {
	s1 := ss.setMap[a]
	s2 := ss.setMap[b]
	if ss.compare(s1.data, s2.data) == true {
		ss.Union(a, b)
	}
}
