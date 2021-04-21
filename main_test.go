package main

import (
	"os"
	"testing"

	"tbd/l3afd/config"
)

func TestTestConfigValid(t *testing.T) {
	// skip if file DNE
	f, err := os.Open("config/l3afd.cfg")
	if err != nil {
		t.Skip("l3afd.cfg not found")
	}
	f.Close()
	if _, err := config.ReadConfig("config/l3afd.cfg"); err != nil {
		t.Errorf("Unable to read l3afd config: %s", err)
	}
}