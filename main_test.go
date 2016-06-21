package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	run := m.Run()
	defer os.Exit(run)
}
