# Clock
> A wrapper around Go's [time](https://golang.org/pkg/time/) to ease testing.

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/transcelestial/clock/Test?label=test&style=flat-square)](https://github.com/transcelestial/clock/actions?query=workflow%3ATest)

The sole purpose of this package is to provide a way to test code that uses time functions from Go's [time](https://golang.org/pkg/time/) package. While there's ways to get around (e.g. make durations for tickers configurable so we can override during tests, etc), it's always better to have more control over time functions like tickers (when the next tick happens), timers (when the timer expires), etc.

**NOTE**: Only `Now()` and `NewTicker()` are implemented as of now.

## Usage

```go
// mypackage.go
package mypackage

import (
    "fmt"
	
	"github.com/transcelestial/clock"
)

func MyFunc(c *clock.Clock) string {
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
	"github.com/transcelestial/clock/clocktest"
)

func TestMyFunc(t *testing.T) {
    // create a time queue from which .Now() will take it's values
    q := clocktest.NewTimequeue(1)
    // and a test clock that uses the queue
    c := clocktest.New(clocktest.WithTimequeue(q))

    // push some values into the queue
    now := time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC)
	q.Push(now)

    assert.Equal(t, "2018-12-31T00:00:00Z", MyFunc(c))
}

```

See [example_clock_test](./example_clock_test.go) and [clocktest/example_clock_test](./clocktest/example_clock_test.go) for more examples.

## TODO
Implement the rest of "time" functions:

* [ ] `time.After`
* [ ] `time.AfterFunc`
* [ ] `time.NewTimer`
* [ ] `time.Sleep`
* [ ] `time.Tick`

## Alternatives
You may also want to try out these alternatives:
* [github.com/benbjohnson/clock](https://github.com/benbjohnson/clock)
* [github.com/thejerf/abtime](https://github.com/thejerf/abtime)

## Contribute
If you wish to contribute, please use the following guidelines:
* Use [conventional commits](https://conventionalcommits.org/)
* Use [effective Go](https://golang.org/doc/effective_go)
