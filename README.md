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

## Use cases

### Send verification email on created a new user

Many web services send an authentication email to new users when they register their email address.
You can use this framework in this situation. An example is shown below.

```go
type OnRegistedUser struct {}

func (o *OnRegisteredUser) Handle(arg interface{}) {
    v, _ := arg.(*AccessObserverArg)
    
    // sending an authentication email
    
    fmt.Printf("sent email to %s(%s)", v.Name, v.Email)
}

type OnRegisteredUserArg struct {
	Email string
	Name string
}

func main() {
    obs := goobserver.NewObservable()
    obs.Attach(&OnRegistedUser{})

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	// register a new user
    	
    	go obs.Notify(&OnRegisteredUserArg{
    	    Email: "...",
    	    Name: "...",
        })
    })
    http.ListenAndServe(":8080", nil)
}
```