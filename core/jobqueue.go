package core

// JobQueue is a producer-consumer queue used to process data items(urls,
// text content for analysis etc...)
type JobQueue interface {
	// Producers queue new data items to be processed.
	Enqueue([]string) (bool, error)

	// Consumers remove data from the queue.
	Dequeue(n int) ([]string, error)
}
