This is an implementation of the union-find datastructure. I'll add more tests and features in the future.
I've used it for help with Coursera Algorithms Part 2 programming assignments and it's done good there.

  <code>go get github.com/Niessy/go-union-find</code>

Short example of using the package:

  
    type Point struct {
        x,y float32
    }
    
    func eucDistance(p1, p2 Point) float32 {
        // .......
    }
    
    func main() {
        // Comparing by euclidian distances
        ss := MakeSetSpace(eucDistance)
        point1 := Point{10.0, 50.0}
        point2 := Point{5.35, 18.75}
        
        // Add our points to our SetSpace
        ss.AddorUpdateSet("first", point1)
        ss.AddorUpdateSet("second", point2)
        
        // Now we can do a union operation
        ss.UnionByCompare("first", "second")
        
        // Note that we can bypass the comparison
        // by using ss.Union("first", "second")
    }
    
    
    
    
