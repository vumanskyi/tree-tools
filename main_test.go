package main

import (
	"bytes"
	"testing"
)

const testTreeFull = `├───lorem
│       └───lorem.txt(9b)
└───project
│       └───test.txt(20b)
`

const testTreeDir = `├───lorem
└───project
`

func TestTreeFull(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "test_data", true)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testTreeFull {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testTreeFull)
	}
}

func TestTreeDir(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "test_data", false)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testTreeDir {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testTreeDir)
	}
}
