package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"log"
)

var DisableControl = false
var userQue UserQue = NewUserQue()

func QueMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isFast := false
		waited := false
		var key = r.Header.Get("Authorization")
		if !DisableControl {
			now := time.Now()
			st := time.Date(now.Year(), now.Month(), now.Day(), 8, 1, 1, 1, time.Local)
			ed := time.Date(now.Year(), now.Month(), now.Day(), 20, 1, 1, 1, time.Local)
			if !(now.After(st) && now.Before(ed)) {
				log.Println("control is off out of date")
				isFast = true //scip control
			}
		}
		if !isFast {
			err := userQue.Wait(key)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				return
			}
			waited = true
		}
		/*TotalProcessingRequestsCounter.Inc()
		log.Info("CounterX:", TotalProcessingRequestsCounter.Value64())*/
		next.ServeHTTP(w, r)
		if waited {
			userQue.Done(key)
		}
		/*TotalProcessingRequestsCounter.Dec()
		log.Info("CounterX:", TotalProcessingRequestsCounter.Value64())*/
	})
}
