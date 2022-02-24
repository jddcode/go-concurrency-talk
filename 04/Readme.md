# Four

Version three used `unbuffered channels` - you can only send information to the channel if the previous message
has already been picked up. Making the first call to `/` on port `8080` worked fine and returned immediately. 
The second call would hang for up to 90 seconds as the agents for `SlowSys` and `MediumSys` processed their first
work packets and eventually picked up the next one from their queues.

In this version we have switched to `buffered channels` - now up to `100` messages can be put into the queue 
before it is full and further add to the queue operations are blocking. Spikes in work can now be stored in the
queue and will be handled in series by the dedicated agents for `SlowSys` and `MediumSys`.

Try making a flurry of requests to `/` on port `8080` and watch as they complete instantly - as long as you make
less than 100 calls at once.

## Code Changes

The new channel used for `SlowSys` is buffered and can store up to 100 messsges

```go
func New() (SlowSys, chan string) {
    channel := make(chan string, 100)
    return slowSys{
        Work: channel,
    }, channel
}
```