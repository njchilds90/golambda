package golambda

// Pipeline provides a fluent, chainable interface over a slice of any type.
// It is intended for use when chaining multiple operations, avoiding
// intermediate variable creation.
//
// Example:
//
//	result := golambda.NewPipeline([]int{1,2,3,4,5,6}).
//	    Filter(func(n int) bool { return n%2 == 0 }).
//	    Map(func(n int) int { return n * 10 }).
//	    Take(2).
//	    Result()
//	// result: [20, 40]
//
// Note: Pipeline[T] is homogeneous — input and output types must match.
// For type-changing transforms (e.g., int → string), use the top-level
// Map function directly.
type Pipeline[T any] struct {
	data []T
}

// NewPipeline creates a new Pipeline from a slice.
// The input slice is copied defensively.
func NewPipeline[T any](data []T) *Pipeline[T] {
	c := make([]T, len(data))
	copy(c, data)
	return &Pipeline[T]{data: c}
}

// Filter applies Filter to the pipeline's data.
func (p *Pipeline[T]) Filter(fn func(T) bool) *Pipeline[T] {
	p.data = Filter(p.data, fn)
	return p
}

// ForEach applies ForEach to the pipeline's data and returns the pipeline unchanged.
func (p *Pipeline[T]) ForEach(fn func(T)) *Pipeline[T] {
	ForEach(p.data, fn)
	return p
}

// Take limits the pipeline to the first n elements.
func (p *Pipeline[T]) Take(n int) *Pipeline[T] {
	p.data = Take(p.data, n)
	return p
}

// Drop removes the first n elements from the pipeline.
func (p *Pipeline[T]) Drop(n int) *Pipeline[T] {
	p.data = Drop(p.data, n)
	return p
}

// TakeWhile keeps elements from the start while fn returns true.
func (p *Pipeline[T]) TakeWhile(fn func(T) bool) *Pipeline[T] {
	p.data = TakeWhile(p.data, fn)
	return p
}

// DropWhile drops elements from the start while fn returns true.
func (p *Pipeline[T]) DropWhile(fn func(T) bool) *Pipeline[T] {
	p.data = DropWhile(p.data, fn)
	return p
}

// Reverse reverses the order of elements in the pipeline.
func (p *Pipeline[T]) Reverse() *Pipeline[T] {
	p.data = Reverse(p.data)
	return p
}

// Result returns the current data slice. The returned slice is a copy.
func (p *Pipeline[T]) Result() []T {
	out := make([]T, len(p.data))
	copy(out, p.data)
	return out
}

// Len returns the number of elements currently in the pipeline.
func (p *Pipeline[T]) Len() int {
	return len(p.data)
}

// Any returns true if fn returns true for any element in the pipeline.
func (p *Pipeline[T]) Any(fn func(T) bool) bool {
	return Any(p.data, fn)
}

// All returns true if fn returns true for all elements in the pipeline.
func (p *Pipeline[T]) All(fn func(T) bool) bool {
	return All(p.data, fn)
}

// None returns true if fn returns true for no elements in the pipeline.
func (p *Pipeline[T]) None(fn func(T) bool) bool {
	return None(p.data, fn)
}
