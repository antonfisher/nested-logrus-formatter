package main

import (
	"fmt"

	"github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	l.Debug("debug message")
	l.Info("info message")
	l.Warn("warn message")
	l.Error("error message")

	ll := l.WithFields(logrus.Fields{
		"component": "main",
	})

	ll.Debug("debug message")
	ll.Info("info message")
	ll.Warn("warn message")
	ll.Error("error message")

	lll := ll.WithFields(logrus.Fields{
		"category": "rest",
	})

	lll.Debug("debug message")
	lll.Info("info message")
	lll.Warn("warn message")
	lll.Error("error message")

	llll := lll.WithFields(logrus.Fields{
		"url": "/load/stats",
	})

	llll.Debug("debug message")
	llll.Info("info message")
	llll.Warn("warn message")
	llll.Error("error message")
	llll.Error(fmt.Errorf("error2"))
	llll.Errorf("error %v", fmt.Errorf("error3"))
}
