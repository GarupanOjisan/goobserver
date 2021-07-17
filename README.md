# goobserver

An observer pattern framework written in Go.

## Usage

1. define a struct implements goobserver.Observer interface
```go
type AccessObserver struct {}

func (o *AccessObserver) Handle(arg interface{}) {
	// do something
}
```

2. Attach your observer

```go
obs := goobserver.NewObservable()
obs.Attach(&AccessObserver{})
```

3. Notify an event to observers

```go
go obs.Notify("access")
```

see more detail in example.