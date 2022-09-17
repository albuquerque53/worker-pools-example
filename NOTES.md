# Problem

Let's suppose that you owner of a book store, and you are moving your store to other address. A bigger place where you'll can expand your business.

In your old book store, you have all your books that need be moved to new store.

You can:
1 - Move book per book to new store.
2 - You can ask for some friends that have cars to carry the books in their cars to move to new place.

What is the best approach?

Obviously, is the second one!

That are worker pools, you have a list of tasks (books) that need be processed (moved) and you have a result (the books present on new address).

# More technicaly...

In worker pools pattern we have tree items:

1 - A list of tasks that must be processed/executed. [Theorical example](https://github.com/albuquerque53/worker-pools-example/blob/main/theorical/theo.go#L13)
2 - A pool of workers that must execute the tasks. [Theorical example]((https://github.com/albuquerque53/worker-pools-example/blob/main/theorical/theo.go#L19)
3 - The final result (the result of each task execution) [Theorical example](https://github.com/albuquerque53/worker-pools-example/blob/main/theorical/theo.go#L16)

So, in Go, the task can be a `Channel` of values (tasks), also the result is a `Channel` that will contain the values of processed tasks.

```go
	tasks := make(chan int, nTasks)
	results := make(chan int, nTasks)
```

We must have a logic (aka: the "processing"), this logic must be isolated in a function, this funcion we call `work` [see this example](https://github.com/albuquerque53/worker-pools-example/blob/main/theorical/theo.go#L37) and must receive the `task`, apply the business logic and produce a `result`.

The responsability of each `worker` is execute the `work` method for the `tasks`.

```go
	for worker := 1; worker < nWorkers; worker++ {
		go work(worker, tasks, results)
	}
```

**In this repo there are three examples:**

## 1 - The technical example (theorical/theo.go)

Link: https://github.com/albuquerque53/worker-pools-example/blob/main/theorical/theo.go

The technical example is a very descritive and simple example about how to implement worker pools, using abstract concepts and a simple logic.

I really recomend you to follow this example firt, will help you organize your ideas about Worker Pools.

If want to run this:
```
make run_theo
```

## 2 - The Books example (without worker pools) (books/without_wp/without_wp.go)

Link: https://github.com/albuquerque53/worker-pools-example/blob/main/books/without_wp/without_wp.go

The simple implementation about the book store example, but following the first alternative (when you move book by book to new store).

If want to run this:
```
# be careful, this can be very slow
make run_wwp
```

## 3 - The Books example (with worker pools) (books/without_wp/without_wp.go)

Link: https://github.com/albuquerque53/worker-pools-example/blob/main/books/using_wp/using_wp.go

This is the implementation about the book store example using worker pools, is a more pratical example using struct, methods and a better defined business logic.

If want to run this:
```
# be careful, this can be very slow
make run_wp
```
