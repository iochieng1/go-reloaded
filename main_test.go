package main

import "testing"

func TestHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestBin(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestUp(t *testing.T) {
	input := "go (up)"
	expected := "GO"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestLow(t *testing.T) {
	input := "STOP SHOUTING (low)"
	expected := "STOP shouting"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestCap(t *testing.T) {
	input := "hello world (cap)"
	expected := "hello World"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestMultiUp(t *testing.T) {
	input := "this is so exciting (up, 2)"
	expected := "this is SO EXCITING"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestPunctuation(t *testing.T) {
	input := "Hello , world !"
	expected := "Hello, world!"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestQuotes(t *testing.T) {
	input := "' hello world '"
	expected := "'hello world'"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func TestArticle(t *testing.T) {
	input := "There is a apple"
	expected := "There is an apple"

	result := processText(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}
