# golambda

[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/golambda.svg)](https://pkg.go.dev/github.com/njchilds90/golambda)
[![CI](https://github.com/njchilds90/golambda/actions/workflows/ci.yml/badge.svg)](https://github.com/njchilds90/golambda/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/njchilds90/golambda)](https://goreportcard.com/report/github.com/njchilds90/golambda)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**Composable, type-safe functional pipeline primitives for Go — using generics.**

Go has always made you write the same `for` loops over and over. `golambda` fixes that with a minimal, zero-dependency, idiomatic library of functional utilities that Python has `itertools`/`toolz`, JavaScript has `lodash`, and Rust has `Iterator` — but Go never had a canonical answer for.

## Install
```bash
go get github.com/njchilds90/golambda@v1.0.0
```

Requires **Go 1.21+**.

## Quick Start
```go
import "github.com/njchilds90/golambda"

nums := []int{1, 2, 3, 4, 5, 6}

evens := golambda.Filter(nums, func(n int) bool { return n%2 == 0 })
// [2, 4, 6]

doubled := golambda.Map(evens, func(n int) int { return n * 2 })
// [4, 8, 12]

sum := golambda.Reduce(doubled, 0, func(acc, n int) int { return acc + n })
// 24
```

### Fluent Pipeline
```go
result := golambda.NewPipeline([]int{1, 2, 3, 4, 5, 6}).
    Filter(func(n int) bool { return n%2 == 0 }).
    Take(2).
    Reverse().
    Result()
// [4, 2]
```

## Function Reference

| Function | Description |
|---|---|
| `Map(s, fn)` | Transform each element |
| `Filter(s, fn)` | Keep matching elements |
| `Reduce(s, init, fn)` | Fold to single value |
| `ForEach(s, fn)` | Side-effectful iteration |
| `Any / All / None` | Predicate checks |
| `Find / FindIndex` | Search for element |
| `GroupBy(s, fn)` | Bucket by key |
| `Partition(s, fn)` | Split into true/false |
| `Chunk(s, n)` | Split into fixed-size groups |
| `Flatten(ss)` | Flatten `[][]T` → `[]T` |
| `FlatMap(s, fn)` | Map then flatten |
| `Zip / Unzip` | Pair and unpair slices |
| `Unique / UniqueBy` | Deduplicate |
| `Contains / IndexOf` | Membership |
| `Reverse` | Reverse order |
| `Take / Drop` | Slice manipulation |
| `TakeWhile / DropWhile` | Predicate-based slicing |
| `CountBy` | Count matching elements |
| `KeyBy` | Index slice as map |
| `Intersect / Difference / Union` | Set operations |
| `Compact` | Remove zero values |
| `MapKeys / MapValues` | Map over map entries |
| `FilterMap / FilterMapErr` | Map + filter in one pass |
| `Tee` | Copy slice n times |
| `SumBy / MinBy / MaxBy` | Aggregations |
| `NewPipeline[T]` | Fluent chainable API |

## Design Principles

- **Zero dependencies** — only the Go standard library
- **Pure functions** — no hidden global state
- **Nil-safe** — nil inputs return nil outputs where appropriate
- **No panics** — all errors are returned explicitly
- **AI-agent friendly** — deterministic, composable, typed errors

## License

MIT © njchilds90
```

---

## Release Instructions (GitHub UI Only)

### Create the v1.0.0 tag and release:

1. In your repo, click **Releases** (right sidebar) → **Create a new release**
2. Click **Choose a tag** → type `v1.0.0` → click **Create new tag: v1.0.0**
3. Target branch: `main`
4. Release title: `v1.0.0`
5. Description: paste from CHANGELOG.md under `[1.0.0]`
6. Click **Publish release**

### Trigger pkg.go.dev indexing:

After publishing the release, visit this URL in your browser to trigger immediate indexing:
```
https://pkg.go.dev/github.com/njchilds90/golambda@v1.0.0
