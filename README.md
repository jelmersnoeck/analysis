# Analysis

Analysis provides a package where you can keep track of you code quality.

## Benchmark

[![GoDoc](https://godoc.org/github.com/jelmersnoeck/analysis/benchmark?status.svg)](https://godoc.org/github.com/jelmersnoeck/analysis/benchmark)

Go has built in Benchmarking support. This is very useful, but running a single
benchmark by itself is useless. You want to know if your benchmark gets better
or - hopefully not - worse over time.

The benchmark package allows to plot this data over time and see if your
functions perform better or worse over time.
