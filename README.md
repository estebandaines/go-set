# go-set
Set implementation for [comparable](https://go.dev/ref/spec#Type_constraints) types.

We propose a basic interface for a Set and its implementation using a map.

To create a new Set:
```
newSet := go-set.NewSet[T]()
```
Type parameter `T` can be replaced with any `comparable` type.

Functions defined for interface `Set`:
```
Add(value T)
Remove(value T)
Contains(value T) bool
Clear()
IsEmpty() bool
Size() int
AsSlice() []T
```

In our implementation, `AsSlice` function do not guaranty order, you should order the result after if required.