# goworker

goworker is implementation of worker-task pattern.


## Description

- WorkerService: it has goroutine(task) pool.
- Task: it's be executed by WorkerService.
- worker: it's runner of Task.

## Usage

```go
// determine worker size
workerSize := 10

// create service
service := goworker.NewService(workerSize)

// generate tasks
for i := 0; i < 100; i++ {
    func(index int) {
        service.Run(func() {
            time.Sleep(time.Second * 1)
            fmt.Printf("Complete Task: %d\n", index)
        })
    }(i)
}
```


## License

MIT
