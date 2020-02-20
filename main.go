package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"time"
	"strconv"

	"github.com/karlpokus/ratelmt"
	"github.com/karlpokus/srv"
)

var Schedule = []int{2, 3, 1}

// index takes a week number and returns the remainder of Schedule len
func Index(week int) int {
	return week % len(Schedule)
}

// the handler returns the schedule for the current week or
// any week if passed by queryparam
func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var weekInt int
		var weekString string
		var err error
		weekString = r.URL.Query().Get("w")
		if weekString != "" {
			weekInt, err = strconv.Atoi(weekString)
			if err != nil {
				http.Error(w, "queryparam w is not a valid int", 400)
				return
			}
		}
		if weekString == "" {
			_, weekInt = time.Now().ISOWeek()
		}
		sched := Schedule[Index(weekInt)]
		date := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "today is:%s\nweek:%d\nschedule:%d", date,  weekInt, sched)
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
