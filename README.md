# Clock
> A wrapper around Go's [time](https://golang.org/pkg/time/) package to ease testing.

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/transcelestial/clock/Test?label=test&style=flat-square)](https://github.com/transcelestial/clock/actions?query=workflow%3ATest)

The sole purpose of this package is to provide a way to test code, using [gomock](https://github.com/golang/mock), that uses time functions from Go's [time](https://golang.org/pkg/time/) package. While there's ways to get around (e.g. make durations for tickers configurable so we can override during tests, etc), it's always better to have more control over time functions like tickers (when the next tick happens), timers (when the timer expires), etc.

**NOTE** Only the following functions are available:
* `Now()`
* `NewTicker()`
* `NewTimer()`
* `Sleep()`

## Usage

```go
// mypackage.go
package mypackage

import (
    "fmt"

    "github.com/transcelestial/clock"
)

func MyFunc(c clock.Clock) string {
    return c.Now().Format(time.RFC3339)
}
```

```go
// mypackage_test.go
package mypackage

import (
    "testing"

    "github.com/transcelestial/clock"

    "github.com/stretchr/testify/assert"
    "github.com/transcelestial/clock/mockclock"
)

func TestMyFunc(t *testing.T) {
    ctrl := gomock.NewController(t)
    // create a mock Clock
    c := mockclock.NewMockClock(ctrl)

    // set some expectations
    now := time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC)
    c.EXPECT().
		Now().
		Return(now)

    assert.Equal(t, "2018-12-31T00:00:00Z", MyFunc(c))
}

```

See [example_clock_test](./example_clock_test.go) and [example_mockclock_test](./example_mockclock_test.go) for more examples.

## TODO
Implement the rest of the "time" functions:

* [ ] `time.After`
* [ ] `time.AfterFunc`
* [x] `time.NewTimer`
* [x] `time.Sleep`
* [ ] `time.Tick`

## Alternatives
You may also want to try out these alternatives:
* [github.com/benbjohnson/clock](https://github.com/benbjohnson/clock)
* [github.com/thejerf/abtime](https://github.com/thejerf/abtime)

## Contribute
If you wish to contribute, please use the following guidelines:
* Use [conventional commits](https://conventionalcommits.org/)
* Use [effective Go](https://golang.org/doc/effective_go)
