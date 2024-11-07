## gfu is go file utils

`GFU` contains following functions:

1. Read file content as slice of strings via `ReadAllLines`
```go
import (
	"github.com/wissance/gfu"
	// other imports
)

var lines []string
var err error

// some operations ...

lines, err = gfu.ReadAllLines("my_file.txt", false)
if len(lines) > 0 {
	
}
// 
```

2. Read file content as a string via `ReadAllText`:
```go
import (
	"github.com/wissance/gfu"
	// other imports
)

var text string
var err error

// some operations ...

text, err = gfu.ReadAllText("my_file.txt")
```
3. Write slice of strings in file via `WriteAllLines`
```go

```
4. Write text in file via `WriteAllText`:
```go

```