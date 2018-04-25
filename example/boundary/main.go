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