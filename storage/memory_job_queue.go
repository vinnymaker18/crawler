// Storage package contains implementations for database/jobqueue related
// interfaces.
package storage

import (
    "utils"
)

// InMemoryJobQueue has everything in memory.
type InMemoryJobQueue struct {
    var jobs []string
    var capLimit int
}

// InMemoryJobQueue constructor
func New(capLimit int) *InMemoryJobQueue {
    return &InMemoryJobQueue{jobs: make([]string, 0, 20), 
        capLimit: capLimit}
}

// JobQueue interface - Enqueue & Dequeue for InMemoryJobQueue
func (queue *InMemoryJobQueue) Enqueue(items []string) (bool, error) {
    if len(queue.jobs) + len(items) > queue.capLimit {
        return false, utils.NewError("Reached max capacity limit.")
    }

    queue.jobs = append(queue.jobs, items...)
    return true, nil
}

func (queue *InMemoryJobQueue) Dequeue(n int) ([]string, error) {
    rem := n
    if rem > len(jobs) {
        rem = len(jobs)
    }

    ret := queue.jobs[:rem]
    queue.jobs = queue.jobs[rem:]
    return ret, nil
}
