package unionfind

import (
	"testing"
)

func intCompare(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func clusterOfInts(n int) []interface{} {
	d := make([]interface{}, n)
	for i := 0; i < n; i++ {
		d[i] = i
	}
	return d
}

type T struct {
	u, v, x int
}

func TestGenerateClustering(t *testing.T) {
	d := []interface{}{1, 2, 3, 4}
	c := GenerateClustering(d, intCompare)
	if c.Count() != len(d) {
		t.Errorf("Clustering should have %d Clusters, instead has %d.\n", len(d), c.Count())
	}
	cMap := make(map[*Cluster]int)
	for _, cluster := range c.AllClusters {
		cMap[cluster] = 0
	}

	if len(cMap) != len(d) {
		t.Errorf("Clustering should have %d leaders, instead has %d.\n", len(d), len(cMap))
	}
}

func TestFindAndUnion(t *testing.T) {
	d := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	c := GenerateClustering(d, intCompare)
	c.UnionbyCompare(c.AllClusters[0], c.AllClusters[1])
	c.UnionbyCompare(c.AllClusters[2], c.AllClusters[3])
	c.UnionbyCompare(c.AllClusters[4], c.AllClusters[5])
	c.UnionbyCompare(c.AllClusters[6], c.AllClusters[7])
	if c.Count() != 4 {
		t.Errorf("Should have 4 leaders, instead has %d.\n", c.Count())
	}
	c.UnionbyCompare(c.AllClusters[1], c.AllClusters[2])
	c.UnionbyCompare(c.AllClusters[5], c.AllClusters[6])
	if c.Count() != 2 {
		t.Errorf("Should have 2 leaders, instead has %d.\n", c.Count())
	}
	c.UnionbyCompare(c.AllClusters[0], c.AllClusters[7])
	if c.Count() != 1 {
		t.Errorf("Should have 1 leaders, instead has %d.\n", c.Count())
	}
}
