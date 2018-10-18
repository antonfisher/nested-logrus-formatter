package main

import (
	"fmt"
	"github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("\n--- nested-logrus-formatter ---\n\n")
	printDemo(&formatter.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category", "req"},
	})

	fmt.Print("\n--- default logrus formatter ---\n\n")
	printDemo(nil)
}

func printDemo(f logrus.Formatter) {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)

	if f != nil {
		l.SetFormatter(f)
	}

	l.Info("this is nested-logrus-formatter demo")

	lWebServer := l.WithField("component", "web-server")
	lWebServer.Info("starting...")

	lWebServerReq := lWebServer.WithFields(logrus.Fields{
		"req":   "GET /api/stats",
		"reqId": "#1",
	})

	lWebServerReq.Info("params: startYear=2048")
	lWebServerReq.Error("response: 400 Bad Request")

	lDbConnector := l.WithField("category", "db-connector")
	lDbConnector.Info("connecting to db on 10.10.10.13...")
	lDbConnector.Warn("connection took 10s")

	l.Info("demo end.")
}
