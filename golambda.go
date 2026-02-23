// Package golambda provides composable, type-safe functional pipeline
// primitives for Go using generics. All functions are pure and stateless,
// making them safe for concurrent use and AI agent composition.
//
// Example usage:
//
//	nums := []int{1, 2, 3, 4, 5, 6}
//	result := golambda.Filter(nums, func(n int) bool { return n%2 == 0 })
//	// result: [2, 4, 6]
//
//	doubled := golambda.Map(result, func(n int) int { return n * 2 })
//	// doubled: [4, 8, 12]
package golambda

// Map applies fn to each element of slice and returns a new slice of results.
//
//	words := []string{"hello", "world"}
//	upper := golambda.Map(words, strings.ToUpper)
//	// upper: ["HELLO", "WORLD"]
func Map[T, U any](slice []T, fn func(T) U) []U {
	if slice == nil {
		return nil
	}
	out := make([]U, len(slice))
	for i, v := range slice {
		out[i] = fn(v)
	}
	return out
}

// Filter returns a new slice containing only elements for which fn returns true.
//
//	nums := []int{1, 2, 3, 4, 5}
//	evens := golambda.Filter(nums, func(n int) bool { return n%2 == 0 })
//	// evens: [2, 4]
func Filter[T any](slice []T, fn func(T) bool) []T {
	if slice == nil {
		return nil
	}
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce reduces slice to a single value by applying fn cumulatively.
// initial is the starting accumulator value.
//
//	nums := []int{1, 2, 3, 4}
//	sum := golambda.Reduce(nums, 0, func(acc, n int) int { return acc + n })
//	// sum: 10
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// ForEach applies fn to each element of slice for side effects.
// Returns nothing; use Map when you need results.
func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

// Any returns true if fn returns true for at least one element.
//
//	golambda.Any([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 }) // false
//	golambda.Any([]int{1, 2, 5}, func(n int) bool { return n%2 == 0 }) // true
func Any[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// All returns true if fn returns true for every element.
//
//	golambda.All([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 }) // true
func All[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// None returns true if fn returns true for no elements.
//
//	golambda.None([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 }) // true
func None[T any](slice []T, fn func(T) bool) bool {
	return !Any(slice, fn)
}

// Find returns the first element for which fn returns true, and a bool
// indicating whether such an element was found.
//
//	val, ok := golambda.Find([]int{1, 2, 3}, func(n int) bool { return n > 1 })
//	// val: 2, ok: true
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element for which fn returns true,
// or -1 if not found.
func FindIndex[T any](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// GroupBy groups elements by a key derived from fn, returning a map.
//
//	words := []string{"ant", "bear", "cat", "deer"}
//	byLen := golambda.GroupBy(words, func(s string) int { return len(s) })
//	// byLen: map[3:["ant","cat"] 4:["bear","deer"]]
func GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, v := range slice {
		k := fn(v)
		out[k] = append(out[k], v)
	}
	return out
}

// Partition splits slice into two slices: the first containing elements for
// which fn returns true, the second for which it returns false.
//
//	evens, odds := golambda.Partition([]int{1,2,3,4,5}, func(n int) bool { return n%2==0 })
//	// evens: [2,4], odds: [1,3,5]
func Partition[T any](slice []T, fn func(T) bool) (trueSlice, falseSlice []T) {
	for _, v := range slice {
		if fn(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return
}

// Chunk splits slice into chunks of the given size.
// The last chunk may be smaller than size if len(slice) is not divisible.
// Returns an error if size < 1.
//
//	chunks, _ := golambda.Chunk([]int{1,2,3,4,5}, 2)
//	// chunks: [[1,2],[3,4],[5]]
func Chunk[T any](slice []T, size int) ([][]T, error) {
	if size < 1 {
		return nil, ErrInvalidChunkSize
	}
	if slice == nil {
		return nil, nil
	}
	out := make([][]T, 0, (len(slice)+size-1)/size)
	for size < len(slice) {
		slice, out = slice[size:], append(out, slice[:size])
	}
	if len(slice) > 0 {
		out = append(out, slice)
	}
	return out, nil
}

// Flatten takes a slice of slices and returns a single flat slice.
//
//	golambda.Flatten([][]int{{1,2},{3},{4,5}}) // [1,2,3,4,5]
func Flatten[T any](slices [][]T) []T {
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	out := make([]T, 0, total)
	for _, s := range slices {
		out = append(out, s...)
	}
	return out
}

// FlatMap maps fn over slice and flattens the result one level.
//
//	golambda.FlatMap([]int{1,2,3}, func(n int) []int { return []int{n, n*10} })
//	// [1,10,2,20,3,30]
func FlatMap[T, U any](slice []T, fn func(T) []U) []U {
	out := make([]U, 0, len(slice))
	for _, v := range slice {
		out = append(out, fn(v)...)
	}
	return out
}

// Zip pairs elements from two slices into a slice of Pair[A,B].
// The result length equals the shorter input slice.
//
//	pairs := golambda.Zip([]int{1,2,3}, []string{"a","b","c"})
//	// [{1,"a"},{2,"b"},{3,"c"}]
func Zip[A, B any](as []A, bs []B) []Pair[A, B] {
	n := len(as)
	if len(bs) < n {
		n = len(bs)
	}
	out := make([]Pair[A, B], n)
	for i := range out {
		out[i] = Pair[A, B]{First: as[i], Second: bs[i]}
	}
	return out
}

// Unzip separates a slice of Pair[A,B] into two slices.
func Unzip[A, B any](pairs []Pair[A, B]) ([]A, []B) {
	as := make([]A, len(pairs))
	bs := make([]B, len(pairs))
	for i, p := range pairs {
		as[i] = p.First
		bs[i] = p.Second
	}
	return as, bs
}

// Unique returns a new slice with duplicate elements removed,
// preserving first-occurrence order. Requires comparable T.
//
//	golambda.Unique([]int{1,2,2,3,1}) // [1,2,3]
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{}, len(slice))
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// UniqueBy returns a new slice with duplicates removed by key fn.
// First occurrence wins.
//
//	type User struct{ ID int; Name string }
//	users := []User{{1,"Alice"},{2,"Bob"},{1,"Alice2"}}
//	golambda.UniqueBy(users, func(u User) int { return u.ID })
//	// [{1,"Alice"},{2,"Bob"}]
func UniqueBy[T any, K comparable](slice []T, fn func(T) K) []T {
	seen := make(map[K]struct{}, len(slice))
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		k := fn(v)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Contains returns true if slice contains target.
//
//	golambda.Contains([]string{"a","b","c"}, "b") // true
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of target in slice,
// or -1 if not found.
func IndexOf[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// Reverse returns a new slice with elements in reversed order.
//
//	golambda.Reverse([]int{1,2,3}) // [3,2,1]
func Reverse[T any](slice []T) []T {
	n := len(slice)
	out := make([]T, n)
	for i, v := range slice {
		out[n-1-i] = v
	}
	return out
}

// Take returns the first n elements of slice.
// If n >= len(slice), the full slice is returned.
//
//	golambda.Take([]int{1,2,3,4,5}, 3) // [1,2,3]
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		out := make([]T, len(slice))
		copy(out, slice)
		return out
	}
	out := make([]T, n)
	copy(out, slice[:n])
	return out
}

// Drop returns a new slice with the first n elements removed.
//
//	golambda.Drop([]int{1,2,3,4,5}, 2) // [3,4,5]
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		out := make([]T, len(slice))
		copy(out, slice)
		return out
	}
	if n >= len(slice) {
		return []T{}
	}
	out := make([]T, len(slice)-n)
	copy(out, slice[n:])
	return out
}

// TakeWhile returns elements from the start of slice while fn returns true.
//
//	golambda.TakeWhile([]int{1,2,3,4,1}, func(n int) bool { return n < 3 })
//	// [1,2]
func TakeWhile[T any](slice []T, fn func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range slice {
		if !fn(v) {
			break
		}
		out = append(out, v)
	}
	return out
}

// DropWhile drops elements from the start of slice while fn returns true,
// then returns the remainder.
//
//	golambda.DropWhile([]int{1,2,3,4,1}, func(n int) bool { return n < 3 })
//	// [3,4,1]
func DropWhile[T any](slice []T, fn func(T) bool) []T {
	for i, v := range slice {
		if !fn(v) {
			out := make([]T, len(slice)-i)
			copy(out, slice[i:])
			return out
		}
	}
	return []T{}
}

// CountBy returns the number of elements for which fn returns true.
func CountBy[T any](slice []T, fn func(T) bool) int {
	count := 0
	for _, v := range slice {
		if fn(v) {
			count++
		}
	}
	return count
}

// KeyBy transforms a slice into a map keyed by fn(element).
// If multiple elements share a key, the last one wins.
//
//	type User struct{ ID int; Name string }
//	m := golambda.KeyBy(users, func(u User) int { return u.ID })
func KeyBy[T any, K comparable](slice []T, fn func(T) K) map[K]T {
	out := make(map[K]T, len(slice))
	for _, v := range slice {
		out[fn(v)] = v
	}
	return out
}

// Intersect returns elements present in both a and b.
// Preserves order of a; requires comparable T.
//
//	golambda.Intersect([]int{1,2,3,4}, []int{2,4,6}) // [2,4]
func Intersect[T comparable](a, b []T) []T {
	set := make(map[T]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	out := make([]T, 0)
	for _, v := range a {
		if _, ok := set[v]; ok {
			out = append(out, v)
		}
	}
	return out
}

// Difference returns elements in a that are not in b.
//
//	golambda.Difference([]int{1,2,3,4}, []int{2,4}) // [1,3]
func Difference[T comparable](a, b []T) []T {
	set := make(map[T]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	out := make([]T, 0)
	for _, v := range a {
		if _, ok := set[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

// Union returns a slice containing all unique elements from both a and b.
// Preserves order; elements from a appear before new elements from b.
//
//	golambda.Union([]int{1,2,3}, []int{2,3,4}) // [1,2,3,4]
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]struct{}, len(a)+len(b))
	out := make([]T, 0, len(a)+len(b))
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	for _, v := range b {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Compact removes zero values from a slice of comparable type.
//
//	golambda.Compact([]string{"a","","b","","c"}) // ["a","b","c"]
//	golambda.Compact([]int{0,1,0,2,3})            // [1,2,3]
func Compact[T comparable](slice []T) []T {
	var zero T
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != zero {
			out = append(out, v)
		}
	}
	return out
}

// Pair holds two values of potentially different types.
type Pair[A, B any] struct {
	First  A
	Second B
}

// NewPair constructs a Pair.
func NewPair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{First: first, Second: second}
}

// MapKeys applies fn to each key in a map, returning a new map.
// If fn produces duplicate keys, later values overwrite earlier ones.
func MapKeys[K1, K2 comparable, V any](m map[K1]V, fn func(K1) K2) map[K2]V {
	out := make(map[K2]V, len(m))
	for k, v := range m {
		out[fn(k)] = v
	}
	return out
}

// MapValues applies fn to each value in a map, returning a new map.
func MapValues[K comparable, V1, V2 any](m map[K]V1, fn func(V1) V2) map[K]V2 {
	out := make(map[K]V2, len(m))
	for k, v := range m {
		out[k] = fn(v)
	}
	return out
}

// FilterMap applies fn to each element; elements for which fn returns
// (value, true) are included in the output. Combines map + filter in one pass.
//
//	strs := []string{"1","two","3","four"}
//	nums, _ := golambda.FilterMap(strs, strconv.Atoi) // [1, 3]  (errors skipped)
//
// Note: fn must return (U, bool) not (U, error). Use FilterMapErr for error variants.
func FilterMap[T, U any](slice []T, fn func(T) (U, bool)) []U {
	out := make([]U, 0, len(slice))
	for _, v := range slice {
		if u, ok := fn(v); ok {
			out = append(out, u)
		}
	}
	return out
}

// FilterMapErr applies fn to each element; elements for which fn returns
// a nil error are included. Non-nil errors are silently skipped.
//
//	strs := []string{"1","two","3"}
//	nums := golambda.FilterMapErr(strs, strconv.Atoi) // [1,3]
func FilterMapErr[T, U any](slice []T, fn func(T) (U, error)) []U {
	out := make([]U, 0, len(slice))
	for _, v := range slice {
		if u, err := fn(v); err == nil {
			out = append(out, u)
		}
	}
	return out
}

// Tee splits a slice into n copies. Useful for applying multiple pipelines
// to the same data without re-iterating.
//
//	a, b := golambda.Tee([]int{1,2,3}, 2)[0], golambda.Tee(...)
func Tee[T any](slice []T, n int) [][]T {
	if n <= 0 {
		return nil
	}
	copies := make([][]T, n)
	for i := range copies {
		c := make([]T, len(slice))
		copy(c, slice)
		copies[i] = c
	}
	return copies
}

// SumBy returns the sum of values produced by fn over slice.
//
//	golambda.SumBy([]int{1,2,3,4}, func(n int) int { return n }) // 10
func SumBy[T any, N Number](slice []T, fn func(T) N) N {
	var total N
	for _, v := range slice {
		total += fn(v)
	}
	return total
}

// MinBy returns the element for which fn produces the smallest value.
// Returns zero value and false if slice is empty.
func MinBy[T any, N Ordered](slice []T, fn func(T) N) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	best := slice[0]
	bestVal := fn(best)
	for _, v := range slice[1:] {
		if val := fn(v); val < bestVal {
			best, bestVal = v, val
		}
	}
	return best, true
}

// MaxBy returns the element for which fn produces the largest value.
// Returns zero value and false if slice is empty.
func MaxBy[T any, N Ordered](slice []T, fn func(T) N) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	best := slice[0]
	bestVal := fn(best)
	for _, v := range slice[1:] {
		if val := fn(v); val > bestVal {
			best, bestVal = v, val
		}
	}
	return best, true
}

// Number is a type constraint for numeric types.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Ordered is a type constraint for types supporting < and >.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
