package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Testing Dictionary Search", func(t *testing.T) {
		dict := map[string]string{"test": "This is a dictionary test"}
		got := Search(dict, "test")
		want := "This is a dictionary test"
		if got != want {
			t.Errorf("got %q, want %q, given %q", got, want, "test")
		}
	})
}
