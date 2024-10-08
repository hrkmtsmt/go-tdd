package hello_test

import (
	"app/hello"
	"testing"
)

func TestHello(t *testing.T) {
	actual := hello.Hello("Jay")
	expect := "Hello Jay!"
	if actual != expect {
		t.Errorf("%v want %v", actual, expect)
	}
}
