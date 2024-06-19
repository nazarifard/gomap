## GoMap
A simple go map wrapper with iterator.

# Example
```go
 import "github.com/nazarifard/gomap"

 func main() {
    m := gomap.New[string, int]()
	
    m.Set("one", 1)
    m.Set("two", 2)
	
    it := m.Iterator()
    for it.Next() {
	print(it.Key(), ": ", it.Value())
    }
 }
 ```


