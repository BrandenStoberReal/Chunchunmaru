package macros

import (
	"bytes"
	"testing"
)

func TestEasy(t *testing.T) {
	tmp, err := BuildTemplate("easy", tmp_easy)
	if err != nil {
		t.Fatalf("Error parsing template: %s", err)
	}
	out := new(bytes.Buffer)
	err = tmp.Execute(out, nil)
	if err != nil {
		t.Errorf("Error executing template: %s", err)
	}
	printTestResults("Easy HTML Template", out.String())
}

func TestMedium(t *testing.T) {
	tmp, err := BuildTemplate("medium", tmp_medium)
	if err != nil {
		t.Fatalf("Error parsing template: %s", err)
	}
	out := new(bytes.Buffer)
	err = tmp.Execute(out, nil)
	if err != nil {
		t.Errorf("Error executing template: %s", err)
	}
	printTestResults("Easy HTML Template", out.String())
}

func TestHard(t *testing.T) {
	tmp, err := BuildTemplate("hard", tmp_hard)
	if err != nil {
		t.Fatalf("Error parsing template: %s", err)
	}
	out := new(bytes.Buffer)
	err = tmp.Execute(out, nil)
	if err != nil {
		t.Errorf("Error executing template: %s", err)
	}
	printTestResults("Easy HTML Template", out.String())
}
