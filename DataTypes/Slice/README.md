# Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```go
a[1:4]
```

## Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

**Changing the elements of a slice modifies the corresponding elements of its underlying array.**

Other slices that share the same underlying array will see those changes.

- Example at [Slice Pointers](SlicePointers.go) (Source : Tour of Go)

## Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

```go
var a [10]int
```
these slice expressions are equivalent:
```go
a[0:10]
a[:10]
a[0:]
a[:]
```