package main

import (
	"strings"
	"testing"
)

func TestReadRecords(t *testing.T) {
	in := `id, name, reporters
0, Markus, "Erik; Artashes"
1, Erik, ""
2, Artashes, ""`

	want := []Record{{"0", " Markus", ` "Erik; Artashes"`},
		{"1", " Erik", ` ""`},
		{"2", " Artashes", ` ""`},
	}

	got, err := ReadRecords(strings.NewReader(in))
	if err != nil {
		t.Errorf("ReadRecords(%q) == Err: %v, want %q", in, err, want)
		return
	}
	for i, _ := range got {
		if strings.Join(got[i], " ") != strings.Join(want[i], " ") {
			t.Errorf("ReadRecords(%q) == \n%q, \nwant %q", in, got[i], want[i])
		}
	}

}
