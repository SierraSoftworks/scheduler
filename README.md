# Scheduler
**An idiomatic Go library providing extensible task scheduling capabilities**

Scheduler provides a framework on which you can quickly and easily schedule
tasks through a series of pluggable scheduling strategies. It aims to offer
a straightforward approach to task scheduling with simple execution cancellation,
error handling and an idiomatic API.

```go
package example

import(
    "log"
    "time"

    "github.com/SierraSoftworks/scheduler"
)

func Example_Schedule() {
    s := scheduler.
        Do(func(t time.Time) error {
            log.Print("Running scheduled task")
        }).
        Every(30 * time.Second).
        Schedule()

    s.CancelWhen(time.After(2 * time.Minute))

    for err := range s.Errors() {
        log.Printf("Failed: %s\n", err.Error())
    }
}
```