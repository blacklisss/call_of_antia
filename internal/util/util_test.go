package util

import (
	"os"
	"testing"
	"time"
)

func TestDirExists(t *testing.T) {
	dirName := "testdir"

	// Cleanup before test
	os.RemoveAll(dirName)

	// Test nonexistent directory
	exists, err := DirExists(dirName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if exists {
		t.Fatalf("Expected directory to not exist")
	}

	// Create a directory
	os.Mkdir(dirName, os.ModePerm)
	defer os.RemoveAll(dirName)

	// Test existing directory
	exists, err = DirExists(dirName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !exists {
		t.Fatalf("Expected directory to exist")
	}
}

func TestMakeDir(t *testing.T) {
	dirName := "testdir_makedir"
	defer os.RemoveAll(dirName)

	err := MakeDir(dirName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = os.Stat(dirName)
	if err != nil {
		t.Fatalf("Expected directory to be created, got %v", err)
	}
}

func TestGetMD5Hash(t *testing.T) {
	text := "Hello, World!"
	expected := "65a8e27d8879283831b664bd8b7f0ad4"
	result := GetMD5Hash(text)

	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestCreateExpireDate(t *testing.T) {
	duration := 5 * time.Second
	expireDate := CreateExpireDate(duration)

	expectedTime := time.Now().Add(duration)

	if !expireDate.Before(expectedTime.Add(1*time.Second)) || !expireDate.After(expectedTime.Add(-1*time.Second)) {
		t.Fatalf("Time mismatch: expected near %v, got %v", expectedTime, expireDate)
	}
}

func TestCompareDate(t *testing.T) {
	date1 := time.Date(2023, 9, 25, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 9, 26, 0, 0, 0, 0, time.UTC)

	if !CompareDate(date1, date2) {
		t.Fatalf("Expected date1 to be before or equal to date2")
	}

	if CompareDate(date2, date1) {
		t.Fatalf("Expected date2 to be after date1")
	}

	date3 := time.Date(2023, 9, 25, 0, 0, 0, 0, time.UTC)
	if !CompareDate(date1, date3) {
		t.Fatalf("Expected date1 to be equal to date3")
	}
}
