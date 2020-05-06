# Maps
- Declare a map of type string-> Vertex (Struct)
```go
var m map[string]Vertex
```
- Make a map for string -> Vertex
```go
m = make(map[string]Vertex)
```
- [Example](Intro.go)

# Mutating Maps
- Insert or update an element in map m:
```go
m[key] = elem
```
- Retrieve an element:

```go
elem = m[key]
```
- Delete an element:

```go
delete(m, key)
```
- Test that a key is present with a two-value assignment:

```go
elem, ok = m[key]
```
- If key is in m, ok is true. If not, ok is false. If key is not in the map, then elem is the zero value for the map's element type.
