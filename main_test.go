package main

import (
  "testing"
  "os"
)

func TestMain(m *testing.M) {
  run := m.Run()
	defer os.Exit(run)
}
