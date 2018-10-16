package formatter_test

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func ExampleFormatter_Format_default() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("test4")

	// Output:
	// - [DEBU] test1
	// - [INFO] test2
	// - [WARN] test3
	// - [ERRO] test4
}

func ExampleFormatter_Format_full_level() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		ShowFullLevel:   true,
	})

	l.Debug("test1")
	l.Info("test2")
	l.Warn("test3")
	l.Error("test4")

	// Output:
	// - [DEBUG] test1
	// - [INFO] test2
	// - [WARNING] test3
	// - [ERROR] test4
}
func ExampleFormatter_Format_show_keys() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		HideKeys:        false,
	})

	ll := l.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")

	// Output:
	// - [INFO] test1
	// - [INFO] [category:rest] test2
}

func ExampleFormatter_Format_hide_keys() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		HideKeys:        true,
	})

	ll := l.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")

	// Output:
	// - [INFO] test1
	// - [INFO] [rest] test2
}

func ExampleFormatter_Format_sort_order() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		HideKeys:        false,
	})

	ll := l.WithField("component", "main")
	lll := ll.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")
	lll.Info("test3")

	// Output:
	// - [INFO] test1
	// - [INFO] [component:main] test2
	// - [INFO] [category:rest] [component:main] test3
}

func ExampleFormatter_Format_field_order() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&formatter.Formatter{
		NoColors:        true,
		TimestampFormat: "-",
		FieldsOrder:     []string{"component", "category"},
		HideKeys:        false,
	})

	ll := l.WithField("component", "main")
	lll := ll.WithField("category", "rest")

	l.Info("test1")
	ll.Info("test2")
	lll.Info("test3")

	// Output:
	// - [INFO] test1
	// - [INFO] [component:main] test2
	// - [INFO] [component:main] [category:rest] test3
}
