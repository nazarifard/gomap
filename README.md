## GoMap
A simple go map wrapper with Range and iter.Seq2 iterator.

# Example
```go
 import "github.com/nazarifard/gomap"

 func main() {
    m := gomap.New[string, int]()	
    m.Set("one", 1)
    m.Set("two", 2)
	m.Range(func(key string, value int) bool {
		print("{", key, ":", value, "}, ")
		return true
	})
	
	for key, value := range m.All() {
		print("{", key, ":", value, "}, ")
	}
 }
 ```


