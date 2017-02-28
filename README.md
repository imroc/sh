sh
==============
sh is a light weight tool for using golang to execute shell script.

## Useage
``` sh
go get github.com/imroc/sh
```
```go
import (
	"fmt"
	"github.com/imroc/sh"
)
func main() {
	s, _ := sh.String("ls -lah") // execute and get string output
	fmt.Println(s)
	sh.Run("./control start")
}
```
