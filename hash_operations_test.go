package main

import "testing"

func TestFile1(t *testing.T) {
	filepath := "./test_files/file.txt"
	want := "a2f113596c6e1278ed7e53e92b707ac0f968050f5a2604b1f2b078fa9d61528c"

	actual, err := hash_file(filepath)
	if err != nil {
		t.Errorf("Hashing function failed unexpectedly: %v", err)
	}

	if actual != want {
		t.Errorf("Hashing function did not produce expected result: %q != %q", actual, want)
	}

}

func TestFile2(t *testing.T) {
	filepath := "./test_files/file2.txt"
	want := "68558e4bcc691d6e84c7e6d3000f9273ce2c237fe874f6ed0ab86013433c500e"

	actual, err := hash_file(filepath)
	if err != nil {
		t.Errorf("Hashing function failed unexpectedly: %v", err)
	}

	if actual != want {
		t.Errorf("Hashing function did not produce expected result: %q != %q", actual, want)
	}

}
