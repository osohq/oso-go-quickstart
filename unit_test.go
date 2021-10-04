package main

import (
   "strings"
   "testing"
   oso "github.com/osohq/go-oso"
)


func TestFailingALot(t *testing.T) {
	var o oso.Oso
	var err error
	if o, err = oso.NewOso(); err != nil {
		t.Fatalf("Failed to set up Oso: %v", err)
	}
 
	o.LoadString("f(x) if x.Foo();")
 
	// Do it 100 times, hoping for bad stuff to happen.
	for i := 0; i < 100; i++ {
		_, errors := o.QueryStr("f(1)")
 
		if err = <-errors; err != nil {
			if !strings.Contains(err.Error(), "'1' object has no attribute 'Foo'") {
				t.Error("Expected Polar runtime error, got none")
			}
		} else {
			t.Fatal("oops")
		}
	}
}

