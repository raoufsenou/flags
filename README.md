flags
=====

it a set of costum defined flags that i need for my own project and which can be easily used by the "flag" package from the standard library.

The available flags are:

## Boundary flags

You can use them like this:

```go
package main

import (
	"flag"
	"log"
	"math"

	"github.com/raoufsenou/flags"
)

func main() {
	fullBoundary := flags.BoundaryValue{Start: 0, End: math.MaxUint32}
	bv := flags.Boundary("bv", fullBoundary, "boundary search limits using flags.Boundary")
	var rv flags.BoundaryValue
	flags.BoundaryVar(&rv, "rv", fullBoundary, "boundary search limits using flags.BoundaryVar")
	flag.Parse()

	log.Printf("type(%T), start:%v, end:%v\n", *bv, bv.Start, bv.End)
	log.Printf("type(%T), start:%v, end:%v\n", rv, rv.Start, rv.End)
}
```

then you can run th file like this:

```
$ go run main.go -bv=end:15,start:10 -rv=1,16

output:

2018/04/25 20:34:29 type(flags.BoundaryValue), start:10, end:15
2018/04/25 20:34:29 type(flags.BoundaryValue), start:1, end:16
```