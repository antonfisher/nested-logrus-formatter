package main

import (
	"github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category", "req"},
	})

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

	l.Fatal("demo end.")
}
