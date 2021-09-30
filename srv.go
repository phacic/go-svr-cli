package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	neglogrus "github.com/meatballhat/negroni-logrus"
	logrus "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func loggerAfter(entry *logrus.Entry, res negroni.ResponseWriter, latency time.Duration, name string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"status":      res.Status(),
		"text_status": http.StatusText(res.Status()),
		"took":        latency.String(),
	})
}

type respWr struct {
	w http.ResponseWriter
}

func (rw *respWr) Header() http.Header {
	return rw.w.Header()
}

func (rw *respWr) WriteHeader(statusCode int) {
	rw.w.WriteHeader(statusCode)
}

func (rw *respWr) Write(data []byte) (int, error) {
	return rw.w.Write(data)
}

func (rw *respWr) setJSONContentHeader() {
	rw.Header().Set("Content-Type", "application/json")
}

func dummyHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw := &respWr{w: res}
	rw.setJSONContentHeader()

	data := make(map[string]interface{})
	data["message"] = "Hello World!"

	dataStr, _ := json.Marshal(data)
	rw.Write(dataStr)
}

// Up handle countup request
func upHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw := &respWr{w: res}
	rw.setJSONContentHeader()

	// stop parameter
	stop := req.URL.Query().Get("stop")

	stopValue, err := strconv.Atoi(stop)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{
			"msg": "stop query param is required and has to be a valid integer"
		}`))
		return
	}

	v, _ := CountUp(stopValue, false)

	retData := make(map[string]interface{})
	retData["counts"] = v
	retData["stop"] = stopValue

	ret, _ := json.Marshal(retData)

	rw.Write(ret)

}

func downHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw := &respWr{w: res}
	rw.setJSONContentHeader()

	start := req.URL.Query().Get("start")

	startValue, err := strconv.Atoi(start)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"msg": "start query param is required and has to be a valid integer"}`))
		return
	}

	v, _ := CountDown(startValue, false)

	retData := make(map[string]interface{})
	retData["start"] = startValue
	retData["counts"] = v

	ret, _ := json.Marshal(retData)
	rw.Write(ret)
}

// RunServer actually run the server which receives
// and processes request
func RunServer() {
	router := httprouter.New()

	// routes
	router.GET("/up", upHandler)
	router.GET("/down", downHandler)
	router.GET("/dummy", dummyHandler)

	neg := negroni.New()

	// logging
	loggerMiddleware := neglogrus.NewMiddlewareFromLogger(Logger, "web")
	loggerMiddleware.After = loggerAfter
	loggerMiddleware.SetLogStarting(true)
	loggerMiddleware.SetLogCompleted(true)

	neg.Use(loggerMiddleware)

	// recovery
	recovery := negroni.NewRecovery()
	recovery.Formatter = &negroni.HTMLPanicFormatter{}
	neg.Use(recovery)

	neg.UseHandler(router)
	neg.Run(":4000")
}
