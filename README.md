## gfu is go file utils

`GFU` contains following functions:

![GFU](/docs/gfu.png)

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
This Write overwrites existing file, to add some lines **_without file erasing_** use `AppendAllLines`
```go
import (
	"github.com/wissance/gfu"
	// other imports
)

lines := []string{
"{",
"    \"id\": 1,",
"    \"name\": \"Michael Ushakov\"",
"}",
}
file := "write_all_lines_test.txt"
err := gfu.WriteAllLines(file, lines, "\n")
```
4. Write text in file via `WriteAllText`, this is _boring_ wrap around `os.WriteFile`
```go
import (
	"github.com/wissance/gfu"
	// other imports
)

text := "some text"
var err error

// some operations ...

err = gfu.WriteAllText("my_file.txt", text)
```

5. Append lines to existing file `AppendAllText`, if file doesn't exist this function  create it 
```go
import (
	"github.com/wissance/gfu"
	// other imports
)

lines := []string{
"{",
"    \"id\": 2,",
"    \"name\": \"Alex Petrov\"",
"}",
}
file := "write_all_lines_test.txt"
err := gfu.AppendAllLines(file, lines, "\n")
```