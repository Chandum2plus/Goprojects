package testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(5, 4)
	want := 9
	if got != want {
		t.Errorf("got %q,wanted %q", got, want)
	}
}
func TestSub(t *testing.T) {
	got := Sub(12, 10)
	want := 2
	if got != want {
		t.Errorf("got %q,wanted %q", got, want)
	}
}
func TestName(t *testing.T) {
	got := Name("Chandu_kumar")
	want := "Chandu_kumar"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSumm(t *testing.T) {
	got := Summ(40, 60)
	want := 100
	if got != want {
		t.Errorf("got %q ,wanted %q", got, want)
	}

}
