package main

import (
	"fmt"
	"testing"
)

func TestEngine(t *testing.T) {
	eng := NewEngine()

	err := eng.recordToFile("mamadus", "sheiparias")

	fmt.Println(eng.lastOffset)

	err = eng.recordToFile("sheetus", "456vfddfbg")

	fmt.Println(eng.lastOffset)

	err = eng.recordToFile("sheeturweerews", "sheiparias123")

	fmt.Println(eng.lastOffset)

	if err != nil {
		t.Errorf("error testing %s", err)
	}
}
