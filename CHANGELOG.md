# Changelog

All notable changes to this project will be documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).
This project uses [Semantic Versioning](https://semver.org/).

## [1.0.0] - 2025-02-23

### Added
- `Map` — transform slice elements
- `Filter` — keep elements matching predicate
- `Reduce` — fold slice to single value
- `ForEach` — side-effectful iteration
- `Any`, `All`, `None` — predicate checks
- `Find`, `FindIndex` — element search
- `GroupBy` — bucket elements by key
- `Partition` — split into true/false slices
- `Chunk` — split into fixed-size slices
- `Flatten` — flatten `[][]T` to `[]T`
- `FlatMap` — map then flatten
- `Zip`, `Unzip` — pair and unpair two slices
- `Unique`, `UniqueBy` — deduplicate
- `Contains`, `IndexOf` — membership checks
- `Reverse` — reverse a slice
- `Take`, `Drop`, `TakeWhile`, `DropWhile` — slice manipulation
- `CountBy` — count matching elements
- `KeyBy` — index slice as map
- `Intersect`, `Difference`, `Union` — set operations
- `Compact` — remove zero values
- `MapKeys`, `MapValues` — map over map keys/values
- `FilterMap`, `FilterMapErr` — combined map+filter
- `Tee` — copy slice n times
- `SumBy`, `MinBy`, `MaxBy` — aggregations
- `Pipeline[T]` — fluent chainable interface
- `Pair[A, B]` — generic tuple type
- `Number`, `Ordered` type constraints
- Full test suite with table-driven tests
- GitHub Actions CI (Go 1.21/1.22/1.23, race detector)
