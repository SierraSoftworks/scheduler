# Scheduler [![Build Status](https://travis-ci.org/SierraSoftworks/scheduler.svg?branch=master)](https://travis-ci.org/SierraSoftworks/scheduler)
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
            return nil
        }).
        Every(30 * time.Second).
        WithHandler(func (errs <-chan error) {
            for err := range errs {
                log.Printf("Failed: %s\n", err.Error())
            }
        })
        Schedule()

    s.CancelWhen(time.After(2 * time.Minute))
}
```