# Two

In our second version of our application, we introduce very simple `goroutines`. This means the client does not
have to wait for calls to `SlowSys` and `MediumSys` to complete and requests appear to complete instantly.

You can test this application by making a call to `/` on port `8080`.

## Code Change

The use of the `go` keyword passes processing of the calls to `SlowSys` and `MediumSys` into separate 
go processes.

```go
func (wk work) Handler(w http.ResponseWriter, r *http.Request) {
    go wk.callSlowSys()
    go wk.callMediumSys()
    w.WriteHeader(http.StatusOK)
}
```