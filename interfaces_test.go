package learn

import "testing"

var ints = Ints{1, 2, 3, 4}

func TestContains_itemIsPresent(t *testing.T) {
	searchFor := 4
	found := ints.Contains(searchFor)
	if !found {
		t.Errorf("Should be found: %v is present in %v", searchFor, ints)
	}
}

func TestContains_itemIsNotPresent(t *testing.T) {
	searchFor := 5
	found := ints.Contains(searchFor)
	if found {
		t.Errorf("Should not be found: %v is not present in %v", searchFor, ints)
	}
}
