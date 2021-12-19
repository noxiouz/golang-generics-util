# golang-generics-util

[![Tests](https://github.com/noxiouz/golang-generics-util/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/noxiouz/golang-generics-util/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/noxiouz/golang-generics-util/branch/main/graph/badge.svg?token=1CQFQI03JL)](https://codecov.io/gh/noxiouz/golang-generics-util)
[![](https://img.shields.io/github/go-mod/go-version/noxiouz/golang-generics-util.svg)](https://img.shields.io/github/go-mod/go-version/noxiouz/golang-generics-util.svg)

Utility library that explores Go generics (1.18)


## xsync

+ [Synchronized[T]](xsync/synchronized.go)
+ [Pool[T]](xsync/pool.go)

## collection
    
+ [Option[T]](collection/option.go)
+ [Map/FlatMap for Option[T]](collection/option.go)
+ [Expected[T, E]](collection/expected.go)
+ [Set[T]](collection/set.go)
+ [Equal[T]([]T, []T)](collection/util.go)
