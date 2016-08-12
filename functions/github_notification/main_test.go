package main

import (
	"fmt"
	"os"
	"testing"
)

// TestMain ...
func TestMain(m *testing.M) {
	setUp()
	ret := m.Run()
	tearDown()
	os.Exit(ret)
}

func setUp() {
	CreateTables()
	fmt.Println("Setup")
}

func tearDown() {
	DeleteTables()
	fmt.Println("tearDown")
}

func TestProceed(t *testing.T) {
	err := Proceed()
	if err != nil {
		t.Errorf("An error occurred. %s", err)
	}
}
