package hello

import "testing"

func TestHello(t *testing.T) {
	want := "V2: Hello!!!"
	if Hello() != want {
		t.Errorf("Hello() !=%s", want)
	}
}
