# Three

In version three we have a problem; `SlowSys` and `MediumSys` are complaining of being bombarded with thousands of
requests per minute. In order to gain control over the rate of calls we are making to these systems, we move from
passing calls to these systems to goroutines, to passing calls to channels and having a dedicated agent make the
calls to each system.

## Code Changes

Packets of work are sent to channels shared with dedicated agents for `SlowSys` and `MediumSys`, meaning that
calls to these systems are now handled in series and not in parallel - no more flooding with requests.

```go
func (wk work) Handler(w http.ResponseWriter, r *http.Request) {
    wk.SlowSysChannel <- "some task"
    wk.MediumSysChannel <- "some task"
    w.WriteHeader(http.StatusOK)
}
```

The dedicated agent for `SlowSys` runs in a goroutine and loops indefinitely, picking up work from it's channel
and then making calls to `SlowSys` - represented here by a one minute sleep.

```go
func (s slowSys) Run() {
    for {
        task := <- s.Work
        fmt.Println("Sending to SlowSys:", task)
        time.Sleep(time.Minute)
    }
}
```

## Important

Note what happens when you open two browser tabs, and make a call to `/` on port `8080`, and then make the 
same request in the second tab only a few seconds later...