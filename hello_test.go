package spider

import "testing"

func TestHello(t *testing.T) {
	want := "Hello!!!"
	if Hello() != want {
		t.Errorf("Hello() !=%s", want)
	}
}
