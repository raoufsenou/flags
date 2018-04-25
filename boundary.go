package flags

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// BoundaryValue struct that holds the definition of a boundary start-end line
type BoundaryValue struct {
	Start, End uint32
}

// toUnit32 helper function to try converting a string number into uint32
func toUint32(s string) (uint32, error) {
	rets, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(rets), nil
}

// toBoundary helper function that try to convert a given string into BoundaryValue variable
func toBoundary(s string) (BoundaryValue, error) {
	var (
		err error
		bv  BoundaryValue
	)
	flag, sK, eK, pK := true, "start:", "end:", ","
	f := func(ss, p string) bool { return strings.Contains(ss, p) }

	// triage function
	trait := func(ts0, ts1, sk, ek string) (bv BoundaryValue, err error) {
		ts0, ts1 = strings.Replace(ts0, sk, "", 1), strings.Replace(ts1, ek, "", 1)
		if bv.Start, err = toUint32(ts0); err != nil {
			return bv, err
		}
		bv.End, err = toUint32(ts1)
		return bv, err
	}

	switch flag {
	case f(s, pK) && f(s, sK) && f(s, eK): // e.g. "start:10,end:15" or "end:15,start:10"
		ts := strings.Split(s, pK)
		if ts0, ts1 := ts[0], ts[1]; f(ts0, sK) && f(ts1, eK) { // ts0="start:10" and ts1="end:15"
			return trait(ts0, ts1, sK, eK)
		} else if ts0, ts1 := ts[0], ts[1]; f(ts0, eK) && f(ts1, sK) { // ts0="end:15" and ts1="start:10"
			return trait(ts1, ts0, sK, eK)
		} else {
			return bv, fmt.Errorf("unsupported format %s", s)
		}
	case f(s, pK) && !f(s, sK) && !f(s, eK): // e.g. "10,15"
		ts := strings.Split(s, pK)
		return trait(ts[0], ts[1], "", "")
	case (f(s, sK) || (f(s, pK) && f(s, sK))) && !f(s, eK): // e.g. "start:10" or "start:10,"
		ts := strings.Replace(s, sK, "", 1)
		ts = strings.Replace(ts, pK, "", 1)
		bv.End = math.MaxUint32
		bv.Start, err = toUint32(ts)
		return bv, err
	case (f(s, eK) || (f(s, pK) && f(s, eK))) && !f(s, sK): // e.g. "end:15" or "end:15,"
		ts := strings.Replace(s, eK, "", 1)
		ts = strings.Replace(ts, pK, "", 1)
		bv.End, err = toUint32(ts)
		return bv, err
	default: // e.g. "10"
		bv.Start, err = toUint32(s)
		bv.End = math.MaxUint32
		return bv, err
	}
}

// Set flag.Value interface implementation
//
// e.g. "start & end"
//		N1.1		s = "start:10,end:15"
//		N1.2		s = "end:15,start:10"
//		N1.3		s = "start:10"		// in this case it will assume that end=math.MaxUint32 (max possible value)
//		N1.4		s = "end:15"		// in this case will assume that start=0
//		N2.1		s = "10,15"			// in this case start=10 and end=15
//		N2.2		s = "10"			// in this case start=10 and assume that end=math.MaxUnit32
func (bv *BoundaryValue) Set(s string) error {
	rv, err := toBoundary(s)
	if err != nil {
		return err
	}
	*bv = rv
	return nil
}

// String flag.Value interface implementation
func (bv BoundaryValue) String() string {
	return fmt.Sprintf("start:%v,end:%v", bv.Start, bv.End)
}

// Boundary defines a boundary flag with specified name, default value, and usage string.
// The return value is the address of BoundaryValue variable that stores the value of the flag.
func Boundary(name string, value BoundaryValue, usage string) *BoundaryValue {
	bv := &BoundaryValue{Start: value.Start, End: value.End}
	flag.Var(bv, name, usage)
	return bv
}

// BoundaryVar defines a boundary flag with specified name, default value, and usage string.
// The argument bv points to an BoundaryValue variable in which to store the value of the flag.
func BoundaryVar(bv *BoundaryValue, name string, value BoundaryValue, usage string) {
	*bv = value
	flag.Var(bv, name, usage)
}
