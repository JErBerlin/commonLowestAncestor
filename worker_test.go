package main

import (
	"strings"
	"testing"
)

func TestFillWorker(t *testing.T) {

	in := Record{"2", "Artashes", "Akshai; Martin"}

	want := Worker{
		"Artashes",
		2,
		[]string{"Akshai", "Martin"},
	}

	got, err := FillWorker(in)

	if err != nil {
		t.Errorf("FillWorker(%q) == Err: %v, want %q", in, err, want)
		return
	}
	if got.id != want.id || got.name != want.name ||
		strings.Join(got.reporters, " ") != strings.Join(want.reporters, " ") {
		t.Errorf("FillWorker(%q) == %q, want %q", in, got, want)
	}
}
