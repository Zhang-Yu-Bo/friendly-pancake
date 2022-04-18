package router

// ref: https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
// ref: https://pjchender.dev/golang/pkg-time-rate/

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/logger"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
	"golang.org/x/time/rate"
)

type mHandle func(http.ResponseWriter, *http.Request)

type connectInfo struct {
	limiter     *rate.Limiter
	lastVisited time.Time
}

var listOfConnections map[string]connectInfo
var rwLock sync.Mutex

const (
	refreshTime    = time.Minute * 5
	defaultTimeout = time.Second * 3
)

func init() {
	listOfConnections = make(map[string]connectInfo)
	go refreshList()
}

func refreshList() {
	for {
		time.Sleep(time.Minute)

		rwLock.Lock()
		for ip, connection := range listOfConnections {
			if time.Since(connection.lastVisited) >= refreshTime {
				delete(listOfConnections, ip)
			}
		}
		rwLock.Unlock()
	}
}

func getConnections(ip string) *rate.Limiter {
	var result connectInfo
	var exist bool

	rwLock.Lock()
	defer rwLock.Unlock()

	if result, exist = listOfConnections[ip]; !exist {
		limiter := rate.NewLimiter(1, 3)
		listOfConnections[ip] = connectInfo{
			limiter:     limiter,
			lastVisited: time.Now(),
		}
		return limiter
	}

	result.lastVisited = time.Now()
	listOfConnections[ip] = result
	return result.limiter
}

func limitRate(mNext mHandle, needWait bool, method utility.ResponseMethod) mHandle {
	return func(w http.ResponseWriter, r *http.Request) {

		var ip string
		if ip = r.Header.Get("X-FORWARDED-FOR"); ip == "" {
			ip = "127.0.0.1"
		}

		limiter := getConnections(ip)
		if needWait {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()
			if err := limiter.Wait(ctx); err != nil {
				logger.ErrorMessage(err)
				responceMsg(w, r, method, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			if !limiter.Allow() {
				logger.ErrorMessage(errors.New("too many requests"))
				responceMsg(w, r, method, http.StatusTooManyRequests, "too many requests")
				return
			}
		}

		mNext(w, r)
	}
}

func responceMsg(
	w http.ResponseWriter,
	r *http.Request,
	method utility.ResponseMethod,
	statusCode int,
	message string) {

	switch method {
	case utility.QRCode:
		utility.ResponseByQRCode(w, statusCode, message)
	case utility.Page:
		utility.ResponseByPage(w, r, statusCode, message)
	default:
		utility.ResponseByJSON(w, statusCode, message)
	}
}
