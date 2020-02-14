package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"time"

	"github.com/karlpokus/ratelmt"
	"github.com/karlpokus/srv"
)

var Schedule = []int{2, 3, 1}

// index takes a week number and returns the remainder of Schedule len
func Index(week int) int {
	return week % len(Schedule)
}

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		_, week := time.Now().ISOWeek()
		i := Index(week)
		fmt.Fprintf(w, "today is week %d and scheduled week %d", week, Schedule[i])
	}
}

func main() {
	stdout := log.New(os.Stdout, "server ", log.Ldate|log.Ltime)
	stderr := log.New(os.Stderr, "server ", log.Ldate|log.Ltime)
	s, err := srv.New(func(s *srv.Server) error {
		router := s.DefaultRouter()
		router.Handle("/", ratelmt.Mw(1, handler()))
		s.Router = router
		s.Logger = stdout
		s.Host = "0.0.0.0"
		s.Port = "9345"
		return nil
	})
	if err != nil {
		stderr.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		stderr.Fatal(err)
	}
	stdout.Println("main exited")
}
