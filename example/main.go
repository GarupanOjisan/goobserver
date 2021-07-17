package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/garupanojisan/goobserver"
)

type AccessObserver struct {
}

func (a *AccessObserver) Handle(arg interface{}) {
	v, ok := arg.(*AccessObserverArg)
	if !ok {
		return
	}

	fmt.Println("trigger")
	time.Sleep(time.Second)
	fmt.Printf("UserID = %s, Message = %s, AccessedAt = %s\n", v.UserID, v.Message, v.AccessedAt.Format(time.RFC3339))
}

type AccessObserverArg struct {
	UserID     string
	Message    string
	AccessedAt time.Time
}

func main() {
	obs := goobserver.NewObservable()
	obs.Attach(&AccessObserver{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// do something
		userID := r.URL.Query().Get("user_id")
		now := time.Now()
		message := r.URL.Query().Get("message")

		// trigger event
		go obs.Notify(&AccessObserverArg{
			UserID:     userID,
			Message:    message,
			AccessedAt: now,
		})

		w.Write([]byte("ok"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
