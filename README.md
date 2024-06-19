## GoMap
A simple go map wrapper with iterator.

# Example
```go
 import (
    "fmt"
    "github.com/nazarifard/gomap"
 )
 func main() {
    m := gomap.New[string, int]()
	
    m.Set("one", 1)
	m.Set("two", 2)
	
    it := m.Iterator()
	for it.Next() {
		fmt.Println(it.Key(), it.Value())
	}

 }
 ```


