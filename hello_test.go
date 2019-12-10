package hello

import "testing"

func TestHello(t *testing.T) {
	want := "V6: Hello!!!"
	if Hello() != want {
		t.Errorf("Hello() !=%s", want)
	}
}
