package go_set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {
	t.Run("When created then create an empty set", func(t *testing.T) {
		assert.Equal(t, &set[string]{}, NewSet[string]())
		assert.Equal(t, &set[int]{}, NewSet[int]())
		assert.Equal(t, &set[int32]{}, NewSet[int32]())
		assert.Equal(t, &set[int64]{}, NewSet[int64]())
		assert.Equal(t, &set[float32]{}, NewSet[float32]())
		assert.Equal(t, &set[float64]{}, NewSet[float64]())
	})
	t.Run("When Add is called then add the elements to the Set", func(t *testing.T) {
		stringSet := make(set[string])
		stringSet.Add("one")
		stringSet.Add("two")
		stringSet.Add("three")
		assert.Equal(t, set[string]{"one": true, "two": true, "three": true}, stringSet)
		intSet := make(set[int])
		intSet.Add(1)
		intSet.Add(2)
		intSet.Add(3)
		assert.Equal(t, set[int]{1: true, 2: true, 3: true}, intSet)
		int32Set := make(set[int32])
		int32Set.Add(1)
		int32Set.Add(2)
		int32Set.Add(3)
		assert.Equal(t, set[int32]{1: true, 2: true, 3: true}, int32Set)
		int64Set := make(set[int64])
		int64Set.Add(1)
		int64Set.Add(2)
		int64Set.Add(3)
		assert.Equal(t, set[int64]{1: true, 2: true, 3: true}, int64Set)
		float32Set := make(set[float32])
		float32Set.Add(1.1)
		float32Set.Add(2.22)
		float32Set.Add(3)
		assert.Equal(t, set[float32]{1.1: true, 2.22: true, 3: true}, float32Set)
		float64Set := make(set[float64])
		float64Set.Add(1.1)
		float64Set.Add(2.22)
		float64Set.Add(3)
		assert.Equal(t, set[float64]{1.1: true, 2.22: true, 3: true}, float64Set)
	})
	t.Run("When Remove is called then add the elements from the Set", func(t *testing.T) {
		stringSet := make(set[string])
		stringSet.Add("one")
		stringSet.Add("two")
		stringSet.Add("three")
		stringSet.Remove("two")
		assert.Equal(t, set[string]{"one": true, "three": true}, stringSet)
		intSet := make(set[int])
		intSet.Add(1)
		intSet.Add(2)
		intSet.Add(3)
		intSet.Remove(2)
		assert.Equal(t, set[int]{1: true, 3: true}, intSet)
		int32Set := make(set[int32])
		int32Set.Add(1)
		int32Set.Add(2)
		int32Set.Add(3)
		int32Set.Remove(2)
		assert.Equal(t, set[int32]{1: true, 3: true}, int32Set)
		int64Set := make(set[int64])
		int64Set.Add(1)
		int64Set.Add(2)
		int64Set.Add(3)
		int64Set.Remove(2)
		assert.Equal(t, set[int64]{1: true, 3: true}, int64Set)
		float32Set := make(set[float32])
		float32Set.Add(1.1)
		float32Set.Add(2.22)
		float32Set.Add(3)
		float32Set.Remove(1.1)
		assert.Equal(t, set[float32]{2.22: true, 3: true}, float32Set)
		float64Set := make(set[float64])
		float64Set.Add(1.1)
		float64Set.Add(2.22)
		float64Set.Add(3)
		float64Set.Remove(1.1)
		assert.Equal(t, set[float64]{2.22: true, 3: true}, float64Set)
	})
	t.Run("When Clear is called then remove all elements from the set", func(t *testing.T) {
		intSet := set[int]{
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
		}
		intSet.Clear()
		assert.Equal(t, set[int]{}, intSet)
	})
	t.Run("When IsEmpty is called and the map is empty then return true", func(t *testing.T) {
		intSet := set[int]{}
		assert.True(t, intSet.IsEmpty())
	})
	t.Run("When IsEmpty is called and the map is not empty but only contains false values then return true", func(t *testing.T) {
		intSet := set[int]{
			1: false,
			2: false,
			3: false,
			4: false,
			5: false,
			6: false,
		}
		assert.True(t, intSet.IsEmpty())
	})
	t.Run("When IsEmpty is called and the map contains true values then return false", func(t *testing.T) {
		intSet := set[int]{
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
		}
		assert.False(t, intSet.IsEmpty())
	})
	t.Run("When Size is called and the map is empty then return 0", func(t *testing.T) {
		intSet := set[int]{}
		assert.Equal(t, 0, intSet.Size())
	})
	t.Run("When Size is called and the map is not empty but only contains false values then return 0", func(t *testing.T) {
		intSet := set[int]{
			1: false,
			2: false,
			3: false,
			4: false,
			5: false,
			6: false,
		}
		assert.Equal(t, 0, intSet.Size())
	})
	t.Run("When Size is called and the map contains true values then return the number of true values", func(t *testing.T) {
		intSet := set[int]{
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
		}
		assert.Equal(t, 4, intSet.Size())
	})
	t.Run("When Contains is called then return true if element is present or false if not", func(t *testing.T) {
		intSet := set[int]{
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
		}
		assert.True(t, intSet.Contains(1))
		assert.False(t, intSet.Contains(2))
		assert.True(t, intSet.Contains(3))
		assert.False(t, intSet.Contains(4))
		assert.True(t, intSet.Contains(5))
		assert.True(t, intSet.Contains(6))
		assert.False(t, intSet.Contains(7))
	})
	t.Run("When AsSlice is called then return the respective slice", func(t *testing.T) {
		intSet := set[int]{
			1: true,
			2: false,
			3: true,
			4: false,
			5: true,
			6: true,
		}
		result := intSet.AsSlice()
		sort.Slice(result, func(i int, j int) bool { return result[i] < result[j] })
		assert.Equal(t, []int{1, 3, 5, 6}, result)
	})
}
