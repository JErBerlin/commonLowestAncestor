package main

import "testing"

func TestFindLCA(t *testing.T) {
	directory := NewDag()
	directory.AddEmployee("Markus", 0)
	directory.AddEmployee("Erik", 1)
	directory.AddEmployee("Artashes", 2)
	directory.AddLink("Markus","Erik")
	directory.AddLink("Markus","Artashes")
	directory.AddEmployee("Akshai", 3)
	directory.AddEmployee("Martin", 4)
	directory.AddLink("Artashes","Akshai")
	directory.AddLink("Artashes","Martin")

	cases := []struct {
		in []string
		want string
	}{
		{[]string{"Akshai", "Martin"}, "Artashes"},
		{[]string{"Martin", "Akshai"}, "Artashes"},
		{[]string{"Artashes", "Martin"}, "Artashes"},
		{[]string{"Artashes", "Erik"}, "Markus"},
		{[]string{"Markus", "Martin"}, "Markus"},
	}
	CEO := "Markus"
	for _, c := range cases {
		got, err := FindLCA(directory, CEO, c.in[0], c.in[1])
		if err != nil {
			t.Errorf("lca(%q) == Err: %v, want %q", c.in, err, c.want)
			return
		}
		if got != c.want {
			t.Errorf("FindLCA(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
