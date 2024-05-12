package slice

import lsync "sync"

// ForEach iterates over elements of collection and invokes iteratee for each element. `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	var wg lsync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			iteratee(_item, _i)
			wg.Done()
		}(item, i)
	}

	wg.Wait()
}
