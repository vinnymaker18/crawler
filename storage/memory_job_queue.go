// Storage package contains implementations for database/jobqueue related
// interfaces.
package storage

type InMemoryJobQueue struct {
}

func (queue *InMemoryJobQueue) Enqueue(items []string) (bool, error) {
}
