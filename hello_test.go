package hello

import "testing"

func TestHello(t *testing.T) {
	want := "V4: Hello!!!"
	if Hello() != want {
		t.Errorf("Hello() !=%s", want)
	}
}
