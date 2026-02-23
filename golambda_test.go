package golambda_test

import (
	"strconv"
	"testing"

	"github.com/njchilds90/golambda"
)

// --- Map ---

func TestMap(t *testing.T) {
	t.Run("doubles integers", func(t *testing.T) {
		got := golambda.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
		want := []int{2, 4, 6}
		assertSliceEqual(t, got, want)
	})
	t.Run("converts to string", func(t *testing.T) {
		got := golambda.Map([]int{1, 2, 3}, strconv.Itoa)
		want := []string{"1", "2", "3"}
		assertSliceEqual(t, got, want)
	})
	t.Run("nil input returns nil", func(t *testing.T) {
		got := golambda.Map[int, int](nil, func(n int) int { return n })
		if got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})
	t.Run("empty input returns empty", func(t *testing.T) {
		got := golambda.Map([]int{}, func(n int) int { return n })
		if len(got) != 0 {
			t.Fatalf("expected empty, got %v", got)
		}
	})
}

// --- Filter ---

func TestFilter(t *testing.T) {
	t.Run("keeps evens", func(t *testing.T) {
		got := golambda.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
		assertSliceEqual(t, got, []int{2, 4})
	})
	t.Run("none match returns empty", func(t *testing.T) {
		got := golambda.Filter([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 })
		if len(got) != 0 {
			t.Fatalf("expected empty, got %v", got)
		}
	})
	t.Run("nil returns nil", func(t *testing.T) {
		if golambda.Filter[int](nil, func(n int) bool { return true }) != nil {
			t.Fatal("expected nil")
		}
	})
}

// --- Reduce ---

func TestReduce(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		got := golambda.Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
		if got != 10 {
			t.Fatalf("expected 10, got %d", got)
		}
	})
	t.Run("product", func(t *testing.T) {
		got := golambda.Reduce([]int{1, 2, 3, 4}, 1, func(acc, n int) int { return acc * n })
		if got != 24 {
			t.Fatalf("expected 24, got %d", got)
		}
	})
	t.Run("empty uses initial", func(t *testing.T) {
		got := golambda.Reduce([]int{}, 42, func(acc, n int) int { return acc + n })
		if got != 42 {
			t.Fatalf("expected 42, got %d", got)
		}
	})
}

// --- Any / All / None ---

func TestPredicates(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	isEven := func(n int) bool { return n%2 == 0 }
	isPos := func(n int) bool { return n > 0 }
	isNeg := func(n int) bool { return n < 0 }

	if !golambda.Any(nums, isEven) {
		t.Error("Any: expected true for even in [1..5]")
	}
	if golambda.Any([]int{1, 3}, isEven) {
		t.Error("Any: expected false for no evens")
	}
	if !golambda.All(nums, isPos) {
		t.Error("All: expected true, all positive")
	}
	if golambda.All(nums, isEven) {
		t.Error("All: expected false, not all even")
	}
	if !golambda.None(nums, isNeg) {
		t.Error("None: expected true, no negatives")
	}
	if golambda.None(nums, isEven) {
		t.Error("None: expected false, some evens exist")
	}
}

// --- Find ---

func TestFind(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		val, ok := golambda.Find([]int{1, 2, 3}, func(n int) bool { return n > 1 })
		if !ok || val != 2 {
			t.Fatalf("expected (2,true), got (%d,%v)", val, ok)
		}
	})
	t.Run("not found", func(t *testing.T) {
		_, ok := golambda.Find([]int{1, 2, 3}, func(n int) bool { return n > 10 })
		if ok {
			t.Fatal("expected not found")
		}
	})
}

func TestFindIndex(t *testing.T) {
	idx := golambda.FindIndex([]int{10, 20, 30}, func(n int) bool { return n == 20 })
	if idx != 1 {
		t.Fatalf("expected 1, got %d", idx)
	}
	if golambda.FindIndex([]int{1, 2}, func(n int) bool { return n == 99 }) != -1 {
		t.Fatal("expected -1")
	}
}

// --- GroupBy ---

func TestGroupBy(t *testing.T) {
	words := []string{"ant", "bear", "cat", "deer"}
	got := golambda.GroupBy(words, func(s string) int { return len(s) })
	if len(got[3]) != 2 || len(got[4]) != 2 {
		t.Fatalf("unexpected grouping: %v", got)
	}
}

// --- Partition ---

func TestPartition(t *testing.T) {
	evens, odds := golambda.Partition([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	assertSliceEqual(t, evens, []int{2, 4})
	assertSliceEqual(t, odds, []int{1, 3, 5})
}

// --- Chunk ---

func TestChunk(t *testing.T) {
	chunks, err := golambda.Chunk([]int{1, 2, 3, 4, 5}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(chunks) != 3 {
		t.Fatalf("expected 3 chunks, got %d", len(chunks))
	}
	assertSliceEqual(t, chunks[2], []int{5})

	_, err = golambda.Chunk([]int{1, 2}, 0)
	if err == nil {
		t.Fatal("expected error for size=0")
	}
}

// --- Flatten / FlatMap ---

func TestFlatten(t *testing.T) {
	got := golambda.Flatten([][]int{{1, 2}, {3}, {4, 5}})
	assertSliceEqual(t, got, []int{1, 2, 3, 4, 5})
}

func TestFlatMap(t *testing.T) {
	got := golambda.FlatMap([]int{1, 2, 3}, func(n int) []int { return []int{n, n * 10} })
	assertSliceEqual(t, got, []int{1, 10, 2, 20, 3, 30})
}

// --- Zip / Unzip ---

func TestZipUnzip(t *testing.T) {
	pairs := golambda.Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
	if len(pairs) != 3 || pairs[0].First != 1 || pairs[0].Second != "a" {
		t.Fatalf("unexpected pairs: %v", pairs)
	}
	as, bs := golambda.Unzip(pairs)
	assertSliceEqual(t, as, []int{1, 2, 3})
	assertSliceEqual(t, bs, []string{"a", "b", "c"})
}

func TestZipShorter(t *testing.T) {
	pairs := golambda.Zip([]int{1, 2, 3}, []string{"a"})
	if len(pairs) != 1 {
		t.Fatalf("expected len 1, got %d", len(pairs))
	}
}

// --- Unique / UniqueBy ---

func TestUnique(t *testing.T) {
	got := golambda.Unique([]int{1, 2, 2, 3, 1})
	assertSliceEqual(t, got, []int{1, 2, 3})
}

func TestUniqueBy(t *testing.T) {
	type item struct{ ID, Val int }
	items := []item{{1, 10}, {2, 20}, {1, 99}}
	got := golambda.UniqueBy(items, func(i item) int { return i.ID })
	if len(got) != 2 || got[0].Val != 10 {
		t.Fatalf("unexpected: %v", got)
	}
}

// --- Contains / IndexOf ---

func TestContains(t *testing.T) {
	if !golambda.Contains([]string{"a", "b", "c"}, "b") {
		t.Error("expected true")
	}
	if golambda.Contains([]string{"a", "b"}, "z") {
		t.Error("expected false")
	}
}

func TestIndexOf(t *testing.T) {
	if golambda.IndexOf([]int{10, 20, 30}, 20) != 1 {
		t.Error("expected 1")
	}
	if golambda.IndexOf([]int{10, 20}, 99) != -1 {
		t.Error("expected -1")
	}
}

// --- Reverse ---

func TestReverse(t *testing.T) {
	got := golambda.Reverse([]int{1, 2, 3})
	assertSliceEqual(t, got, []int{3, 2, 1})
	// original unmodified
	orig := []int{1, 2, 3}
	golambda.Reverse(orig)
	if orig[0] != 1 {
		t.Error("Reverse should not mutate original")
	}
}

// --- Take / Drop / TakeWhile / DropWhile ---

func TestTakeDrop(t *testing.T) {
	assertSliceEqual(t, golambda.Take([]int{1, 2, 3, 4, 5}, 3), []int{1, 2, 3})
	assertSliceEqual(t, golambda.Take([]int{1, 2}, 10), []int{1, 2})
	assertSliceEqual(t, golambda.Take([]int{1, 2, 3}, 0), []int{})
	assertSliceEqual(t, golambda.Drop([]int{1, 2, 3, 4, 5}, 2), []int{3, 4, 5})
	assertSliceEqual(t, golambda.Drop([]int{1, 2}, 10), []int{})
}

func TestTakeWhileDropWhile(t *testing.T) {
	lt3 := func(n int) bool { return n < 3 }
	assertSliceEqual(t, golambda.TakeWhile([]int{1, 2, 3, 4, 1}, lt3), []int{1, 2})
	assertSliceEqual(t, golambda.DropWhile([]int{1, 2, 3, 4, 1}, lt3), []int{3, 4, 1})
}

// --- CountBy ---

func TestCountBy(t *testing.T) {
	n := golambda.CountBy([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if n != 2 {
		t.Fatalf("expected 2, got %d", n)
	}
}

// --- KeyBy ---

func TestKeyBy(t *testing.T) {
	type user struct{ ID int }
	users := []user{{1}, {2}, {3}}
	m := golambda.KeyBy(users, func(u user) int { return u.ID })
	if m[2].ID != 2 {
		t.Fatal("expected user with ID=2")
	}
}

// --- Set operations ---

func TestSetOps(t *testing.T) {
	assertSliceEqual(t, golambda.Intersect([]int{1, 2, 3, 4}, []int{2, 4, 6}), []int{2, 4})
	assertSliceEqual(t, golambda.Difference([]int{1, 2, 3, 4}, []int{2, 4}), []int{1, 3})
	assertSliceEqual(t, golambda.Union([]int{1, 2, 3}, []int{2, 3, 4}), []int{1, 2, 3, 4})
}

// --- Compact ---

func TestCompact(t *testing.T) {
	assertSliceEqual(t, golambda.Compact([]int{0, 1, 0, 2, 0}), []int{1, 2})
	assertSliceEqual(t, golambda.Compact([]string{"a", "", "b", ""}), []string{"a", "b"})
}

// --- MapKeys / MapValues ---

func TestMapKeysValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	doubled := golambda.MapValues(m, func(v int) int { return v * 2 })
	if doubled["a"] != 2 || doubled["b"] != 4 {
		t.Fatalf("unexpected: %v", doubled)
	}
}

// --- FilterMap / FilterMapErr ---

func TestFilterMap(t *testing.T) {
	got := golambda.FilterMap([]int{1, 2, 3, 4}, func(n int) (int, bool) {
		return n * 2, n%2 == 0
	})
	assertSliceEqual(t, got, []int{4, 8})
}

func TestFilterMapErr(t *testing.T) {
	strs := []string{"1", "two", "3"}
	got := golambda.FilterMapErr(strs, strconv.Atoi)
	assertSliceEqual(t, got, []int{1, 3})
}

// --- SumBy / MinBy / MaxBy ---

func TestSumBy(t *testing.T) {
	got := golambda.SumBy([]int{1, 2, 3, 4}, func(n int) int { return n })
	if got != 10 {
		t.Fatalf("expected 10, got %d", got)
	}
}

func TestMinMaxBy(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9}
	min, ok := golambda.MinBy(nums, func(n int) int { return n })
	if !ok || min != 1 {
		t.Fatalf("min: expected (1,true), got (%d,%v)", min, ok)
	}
	max, ok := golambda.MaxBy(nums, func(n int) int { return n })
	if !ok || max != 9 {
		t.Fatalf("max: expected (9,true), got (%d,%v)", max, ok)
	}
	_, ok = golambda.MinBy([]int{}, func(n int) int { return n })
	if ok {
		t.Fatal("expected false for empty")
	}
}

// --- Tee ---

func TestTee(t *testing.T) {
	copies := golambda.Tee([]int{1, 2, 3}, 3)
	if len(copies) != 3 {
		t.Fatalf("expected 3 copies, got %d", len(copies))
	}
	copies[0][0] = 99
	if copies[1][0] == 99 {
		t.Fatal("Tee copies should be independent")
	}
}

// --- Pipeline ---

func TestPipeline(t *testing.T) {
	result := golambda.NewPipeline([]int{1, 2, 3, 4, 5, 6}).
		Filter(func(n int) bool { return n%2 == 0 }).
		Take(2).
		Reverse().
		Result()
	assertSliceEqual(t, result, []int{4, 2})
}

func TestPipelineLen(t *testing.T) {
	p := golambda.NewPipeline([]int{1, 2, 3, 4})
	if p.Len() != 4 {
		t.Fatalf("expected 4, got %d", p.Len())
	}
}

func TestPipelinePredicates(t *testing.T) {
	p := golambda.NewPipeline([]int{2, 4, 6})
	if !p.All(func(n int) bool { return n%2 == 0 }) {
		t.Error("All evens expected")
	}
	if p.Any(func(n int) bool { return n%2 != 0 }) {
		t.Error("No odds expected")
	}
}

// --- helpers ---

func assertSliceEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %v (len %d), want %v (len %d)", got, len(got), want, len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("index %d: got %v, want %v", i, got[i], want[i])
		}
	}
}
