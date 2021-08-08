package idag

import (
	"reflect"
	"testing"
)

func TestGraph_CAB(t *testing.T) {
	var g Graph

	wantVertex, a := uint(0), g.Append([]uint{})
	if a != wantVertex {
		t.Errorf("want %v, got %v", wantVertex, a)
	}

	wantVertex, b := uint(1), g.Append([]uint{})
	if b != wantVertex {
		t.Errorf("want %v, got %v", wantVertex, b)
	}

	wantVertex, c := uint(2), g.Append([]uint{a, b})
	if c != wantVertex {
		t.Errorf("want %v, got %v", wantVertex, c)
	}

	wantVertices, cAdjacent := []uint{a, b}, g.Read(c)
	if !reflect.DeepEqual(cAdjacent, wantVertices) {
		t.Errorf("want %v, got %v", wantVertices, cAdjacent)
	}
}
