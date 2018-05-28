# strategy
`import "github.com/carousell/Orion/utils/httptripper/strategy"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package strategy provides strategies for use with retry

## <a name="pkg-imports">Imported Packages</a>

No packages beyond the Go standard library are imported.

## <a name="pkg-index">Index</a>
* [type Strategy](#Strategy)
  * [func DefaultStrategy(duration time.Duration) Strategy](#DefaultStrategy)
  * [func ExponentialStrategy(duration time.Duration) Strategy](#ExponentialStrategy)

#### <a name="pkg-files">Package files</a>
[strategy.go](./strategy.go) [types.go](./types.go) 

## <a name="Strategy">type</a> [Strategy](./types.go#L9-L12)
``` go
type Strategy interface {
    //WaitDuration takes attempt, maxRetry and request/response paramaetrs as input and gives out a duration as response
    WaitDuration(attempt, maxRetry int, req *http.Request, resp *http.Response, err error) time.Duration
}
```
Strategy is the interface requirement for any strategy implementation

### <a name="DefaultStrategy">func</a> [DefaultStrategy](./strategy.go#L29)
``` go
func DefaultStrategy(duration time.Duration) Strategy
```
DefaultStrategy provides implementation for Fixed duration wait

### <a name="ExponentialStrategy">func</a> [ExponentialStrategy](./strategy.go#L37)
``` go
func ExponentialStrategy(duration time.Duration) Strategy
```
ExponentialStrategy provided implementation for exponentially (in powers of 2) growing wait duration

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)