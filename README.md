# nested-logrus-formatter

[![Build Status](https://travis-ci.org/antonfisher/nested-logrus-formatter.svg?branch=master)](https://travis-ci.org/antonfisher/nested-logrus-formatter)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonfisher/nested-logrus-formatter)](https://goreportcard.com/report/github.com/antonfisher/nested-logrus-formatter)
[![GoDoc](https://godoc.org/github.com/antonfisher/nested-logrus-formatter?status.svg)](https://godoc.org/github.com/antonfisher/nested-logrus-formatter)

Human-readable log formatter, converts _logrus_ fields to a nested structure:

![Screenshot](https://raw.githubusercontent.com/antonfisher/nested-logrus-formatter/docs/images/demo.png)

## Configuration:

```go
type Formatter struct {
	FieldsOrder     []string // by default fields are sorted alphabetically
	TimestampFormat string   // by default time.StampMilli = "Jan _2 15:04:05.000" is used
	HideKeys        bool     // to show only [fieldValue] instead of [fieldKey:fieldValue]
	NoColors        bool     // to disable all colors
	NoFieldsColors  bool     // to disable colors only on fields and keep levels colored
	ShowFullLevel   bool     // to show full level (e.g. [WARNING] instead of [WARN])
	TrimMessages    bool     // to trim whitespace on messages
}
```

## Usage

```go
import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

log := logrus.New()
log.SetFormatter(&nested.Formatter{
	HideKeys:    true,
	FieldsOrder: []string{"component", "category"},
})

log.Info("just info message")
// Output: Jan _2 15:04:05.000 [INFO] just info message

log.WithField("component", "rest").Warn("warn message")
// Output: Jan _2 15:04:05.000 [WARN] [rest] warn message
```

## Development

```bash
# run tests:
make test

# run demo:
make demo
```
