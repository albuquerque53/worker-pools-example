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

1 - A list of tasks that must be processed/executed.
2 - A pool of workers that must execute the tasks.
3 - The final result (the result of each task execution)

So, in Go, the task can be a `Channel` of values (tasks), also the result is a `Channel` that will contain the values of processed tasks.