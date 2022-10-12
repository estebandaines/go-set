package go_set

type Set[T comparable] interface {
	Add(value T)
	Remove(value T)
	Clear()
	IsEmpty() bool
	Size() int
	Contains(value T) bool
	AsSlice() []T
}

type set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	s := make(set[T])
	return &s
}

func (s *set[T]) Add(value T) {
	(*s)[value] = true
}

func (s *set[T]) Remove(value T) {
	_, ok := (*s)[value]
	if ok {
		delete(*s, value)
	}
}

func (s *set[T]) Contains(value T) bool {
	result, ok := (*s)[value]
	return ok && result
}

func (s *set[T]) Clear() {
	for key := range *s {
		delete(*s, key)
	}
}

func (s *set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *set[T]) Size() int {
	size := 0
	for _, v := range *s {
		if v {
			size++
		}
	}
	return size
}

func (s *set[T]) AsSlice() []T {
	result := make([]T, 0)
	for k, v := range *s {
		if v {
			result = append(result, k)
		}
	}
	return result
}
