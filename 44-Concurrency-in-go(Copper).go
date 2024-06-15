/*

type Future[T any] interface {
    // Await waits for the result
    Await() T
}

type TaskQueueExecutor[T any] interface {
    // Submit returns immediately if there's a space in the task queue, otherwise blocks
    Submit(func() T) Future[T]
}

func NewTaskQueueExecutor[T any](queueSize, poolSize int) TaskQueueExecutor[T] {
    panic("implement me")
}

*/

/*
Implement a TaskQueueExecutor using Go generics, along with a Future interface for asynchronous task handling. 

Future[T any]: An interface to handle the result of asynchronous operations. This should have:

  Await() T: A method that blocks until the result of the future is available and then returns it.

TaskQueueExecutor[T any]: An interface to manage tasks asynchronously. This should have:

  Submit(func() T) Future[T]: A method that submits a task to be executed asynchronously. It returns immediately with a Future, 
  which can later be used to retrieve the result of the task. 
  If the task queue is full, the behavior (whether to block, error out, or drop the task) might need to be specified.

*/

/*
Implementation Strategy

Future Implementation:

Use a channel to communicate the result from the executing goroutine to the caller.
Await() would simply read from this channel, blocking until a value is available.

TaskQueueExecutor Implementation:

Manage a queue of tasks using a channel.
Use a pool of goroutines to execute tasks from the queue.

NewTaskQueueExecutor Function:

Initialize the task queue and start a specified number of worker goroutines.
*/

package main

import (
    "fmt"
    "time"
)

type Future[T any] struct {
    result chan T
}

func (f *Future[T]) Await() T {
    return <-f.result
}

type TaskQueueExecutor[T any] struct {
    tasks chan func() T
}

func NewTaskQueueExecutor[T any](queueSize, poolSize int) *TaskQueueExecutor[T] {
    executor := &TaskQueueExecutor[T]{
        tasks: make(chan func() T, queueSize),
    }

    for i := 0; i < poolSize; i++ {
        go func() {
            for task := range executor.tasks {
                result := task()
                future := &Future[T]{result: make(chan T, 1)}
                future.result <- result
            }
        }()
    }

    return executor
}

func (executor *TaskQueueExecutor[T]) Submit(task func() T) *Future[T] {
    future := &Future[T]{result: make(chan T, 1)}
    go func() {
        executor.tasks <- task
        result := task()
        future.result <- result
    }()
    return future
}

func main() {
    executor := NewTaskQueueExecutor[int](5, 3)
    future := executor.Submit(func() int {
        time.Sleep(2 * time.Second) // simulate work
        return 42
    })
    result := future.Await()
    fmt.Println("Result of task:", result)
}


/*
package main

import (
    "fmt"
    "sync"
    "time"
)

// Future interface
type Future[T any] interface {
    Await() (T, error)
}

// TaskQueueExecutor interface
type TaskQueueExecutor[T any] interface {
    Submit(task func() (T, error)) Future[T]
}

// concrete implementation of Future
type concreteFuture[T any] struct {
    result chan T
    err    chan error
    wg     *sync.WaitGroup
}

func (f *concreteFuture[T]) Await() (T, error) {
    f.wg.Wait() // Ensure the task is complete
    select {
    case res := <-f.result:
        return res, nil
    case err := <-f.err:
        var zero T
        return zero, err
    }
}

// concrete implementation of TaskQueueExecutor
type taskQueueExecutor[T any] struct {
    queue   chan func() (T, error)
    wg      sync.WaitGroup
}

func NewTaskQueueExecutor[T any](queueSize, poolSize int) TaskQueueExecutor[T] {
    executor := &taskQueueExecutor[T]{
        queue: make(chan func() (T, error), queueSize),
    }

    // Start worker goroutines
    for i := 0; i < poolSize; i++ {
        go func() {
            for task := range executor.queue {
                executor.wg.Add(1)
                result, err := task()
                f := &concreteFuture[T]{
                    result: make(chan T, 1),
                    err:    make(chan error, 1),
                    wg:     &executor.wg,
                }
                if err != nil {
                    f.err <- err
                } else {
                    f.result <- result
                }
                executor.wg.Done()
            }
        }()
    }

    return executor
}

func (ex *taskQueueExecutor[T]) Submit(task func() (T, error)) Future[T] {
    f := &concreteFuture[T]{
        result: make(chan T, 1),
        err:    make(chan error, 1),
        wg:     new(sync.WaitGroup),
    }
    f.wg.Add(1)
    go func() {
        defer f.wg.Done()
        ex.queue <- task
    }()
    return f
}

func main() {
    executor := NewTaskQueueExecutor[int](5, 3)
    future := executor.Submit(func() (int, error) {
        time.Sleep(2 * time.Second) // simulate work
        return 42, nil
    })
    result, err := future.Await()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result of task:", result)
    }
}


*/
