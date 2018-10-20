# nested-logrus-formatter

[![Build Status](https://travis-ci.org/antonfisher/nested-logrus-formatter.svg?branch=master)](https://travis-ci.org/antonfisher/nested-logrus-formatter)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonfisher/nested-logrus-formatter)](https://goreportcard.com/report/github.com/antonfisher/nested-logrus-formatter)

Human-readable log formatter, converts _logrus_ fields to a nested structure:

![Screenshot](https://raw.githubusercontent.com/antonfisher/nested-logrus-formatter/docs/images/demo.png)

## Configuration:

```go
type Formatter struct {
	FieldsOrder     []string // default: fields sorted alphabetically
	TimestampFormat string   // default: time.StampMilli = "Jan _2 15:04:05.000"
	HideKeys        bool     // show [fieldValue] instead of [fieldKey:fieldValue]
	NoColors        bool     // disable colors
	NoFieldsColors  bool     // color only level, default is level + fields
	ShowFullLevel   bool     // true to show full level [WARNING] instead [WARN]
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
