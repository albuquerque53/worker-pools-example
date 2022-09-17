# Worker Pools Example

## What are Worker Pools?

Worker pools are a [Concurrency Pattern](https://en.wikipedia.org/wiki/Concurrency_pattern) where we provide two or more "workers" or "executors" to execute multiple "tasks" or "jobs" concurrently.

You can read my notes with more descritive explanation while follow my examples and after that, you can run the benchmarks to see the efficiency of each example

## How to Run (Benchmarks)

The intention of this project is not run the code like some API application os something like that.

There are benchmarks to each example, so that you can compare the execution time of every approach.

I wrote a Makefile to help you run these benchmarks, easy peasy.


1 - To run the books example using worker pools:
```
make run_wp
```

2 - To run the books example without worker pools:
```
# be careful, this can be very slow
make run_wwp
```

3 - To run the theorical example:
```
make run_theo
```

## This repo structure

This repo follows the structure:

```
├── books
│   ├── using_wp
│   │   ├── using_wp.go
│   │   └── using_wp_test.go
│   └── without_wp
│       ├── without_wp.go
│       └── without_wp_test.go
├── theorical
│   └── theo.go
├── go.mod
├── Makefile
├── NOTES.md
└── README.md
```