package main

import (
	"flag"
	"log"
	"math"

	"github.com/raoufsenou/flags"
)

func main() {
	fullBoundary := flags.BoundaryValue{Start: 0, End: math.MaxUint32}
	bv := flags.Boundary("bv", fullBoundary, "boundary search limits")
	var rv flags.BoundaryValue
	flags.BoundaryVar(&rv, "rv", fullBoundary, "hhh")
	flag.Parse()

	log.Printf("%T, %v, %v\n", *bv, bv.Start, bv.End)
	log.Printf("%T, %v, %v\n", rv, rv.Start, rv.End)
}
