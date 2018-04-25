package flags

import (
	"math"
	"testing"
)

func TestBoundary(t *testing.T) {
	var max uint32 = math.MaxUint32
	tests := []struct {
		text        string
		startparsed uint32
		endparsed   uint32
		invalid     bool
	}{
		{text: "end:10", endparsed: 10},
		{text: "end:10,", endparsed: 10},
		{text: "10", startparsed: 10, endparsed: max},
		{text: "10,11", startparsed: 10, endparsed: 11},
		{text: "start:10", startparsed: 10, endparsed: max},
		{text: "start:10,", startparsed: 10, endparsed: max},
		{text: "start:10,end:11", startparsed: 10, endparsed: 11},
		{text: "end:11,start:10", startparsed: 10, endparsed: 11},
		// must fail
		{text: "-1", invalid: true},
		{text: "10,", invalid: true},
		{text: "p:12", invalid: true},
		{text: "-1,10", invalid: true},
		{text: "10,-1", invalid: true},
		{text: "-1,-10", invalid: true},
		{text: "end:-1", invalid: true},
		{text: "start:-1", invalid: true},
		{text: "start:-1,end:11", invalid: true},
		{text: "end:-1,start:10", invalid: true},
		{text: "start:-1,end:-2", invalid: true},
		{text: "end:-1,start:-3", invalid: true},
	}

	for _, tt := range tests {
		bv := &BoundaryValue{}
		if err := bv.Set(tt.text); err != nil {
			if !tt.invalid {
				t.Errorf("parsing %s failed unexpectedly: %v", tt.text, err)
			}
			continue
		}

		if tt.invalid {
			t.Errorf("parsing %s should have failed", tt.text)
			continue
		}

		wants := BoundaryValue{Start: tt.startparsed, End: tt.endparsed}

		if *bv != wants {
			t.Errorf("%s should be parsed as %v; got %v", tt.text, wants, bv)
		}

	}
}
