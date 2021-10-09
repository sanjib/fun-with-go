package main

import "testing"

var tests = []struct {
	name            string
	dividend        float32
	divisor         float32
	expected        float32
	shouldHaveError bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivideStructBased(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.shouldHaveError {
			if err == nil {
				t.Error("expected an error, did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one")
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

//func TestDivide(t *testing.T) {
//	_, err := divide(10.0, 2.0)
//	if err != nil {
//		t.Error("err test:", err)
//	}
//}
//
//func TestBadDivide(t *testing.T) {
//	_, err := divide(10.0, 0)
//	if err == nil {
//		t.Error("did not get err for 10.0/0")
//	}
//}
