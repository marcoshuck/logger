# Logger

## Installation
```
$ go get -u github.com/marcoshuck/logger
```

## Usage

### In your code
```go
func main() {
    l := logger.NewLogger(logger.VerbosityDebug, os.Stdout)
    msg := "Hello world!"
    l.Debug(fmt.Sprintf("Important message: %s", msg))
}
```

### Output

```text
[31 Aug 20 18:34 -0300] [DEBUG] Important message: Hello world!
```


## Verbosity levels
```go
const (
	VerbosityDebug = 0
	VerbosityInfo = 1
	VerbosityWarning = 2
	VerbosityError = 3
	VerbosityCritical = 4
)
```