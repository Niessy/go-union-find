// This is a basic implementation
// of the Union Find Data Structure.
package unionfind

// Minimal Cluster structure
type Cluster struct {
	leader *Cluster    // pointer to leader Cluster
	rank   int         // rank of the cluster
	data   interface{} // whatever you're trying to compute with 
}

// Create a Cluster given data, this can be any type.
func NewCluster(d interface{}) *Cluster {
	nc := new(Cluster)
	nc.rank = 0
	nc.leader = nc
	nc.data = d
	return nc
}

// A slice of Cluster structs
type ClusterGroup []*Cluster

// Returns length of ClusterGroup
func (c ClusterGroup) Len() int {
	return len(c)
}

type compareFunc func(a, b interface{}) bool

// Contains ClusterGroup, count (number of Clusters) and
// the comparison function for the data in the Clusters.
type Clustering struct {
	AllClusters ClusterGroup
	count       int         // count of current seperated Clusters
	compare     compareFunc // function to compare Data
}

// Generates a Clustering with a Clusters 
// based on the passed Data and compareFunc
func GenerateClustering(data []interface{}, cf compareFunc) *Clustering {
	clusterGroup := make(ClusterGroup, len(data))
	for i, d := range data {
		nc := NewCluster(d)
		clusterGroup[i] = nc
	}
	clustering := &Clustering{
		AllClusters: clusterGroup,
		count:       clusterGroup.Len(),
		compare:     cf,
	}
	return clustering
}

//Returns the number of Clusters
func (c Clustering) Count() int {
	return c.count
}

// Find the leader of a Cluster
// Uses path compression so the Operation runs
// in amortized O(alpha(n)) where alpha is 
// the inverse Ackermann function
func (clus Clustering) Find(c *Cluster) *Cluster {
	if c.leader != c {
		c.leader = clus.Find(c.leader)
	}
	return c.leader
}

// Unions two Cluster structs, the Cluster with the highest rank
// is chosen for leadership.
func (clus *Clustering) Union(c, d *Cluster) {
	// find the leaders of c and d clusters
	c1 := *clus
	x1 := c1.Find(c)
	x2 := c1.Find(d)
	// Already in the same cluster
	if x2 == x1 {
		return
	}
	a1 := *x1
	a2 := *x2

	// Merge based on rank of Cluster
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
	*clus = c1
	return
}

// Takes two Cluster structs, works like Union except that
// the predefined compareFunc is used to determine if 
// the two Cluster should Union.
func (clus *Clustering) UnionbyCompare(c, d *Cluster) {
	if clus.compare(c.data, d.data) == true {
		clus.Union(c, d)
	}
}
