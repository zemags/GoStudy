package main

import "testing"

func TestMap(t *testing.T) {

	checkColors := map[string]string{}

	checkColors = addColor(checkColors)
	_, isKeyPresent := checkColors["yellow"]
	if !isKeyPresent { // ! means 'not'
		t.Errorf("There is no yellow in map")
	}
}
