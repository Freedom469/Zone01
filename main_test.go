package main

import(
	"testing"
	"strings"
)

func TestQuotationMarkFunctionality(t *testing.T) {
	Quotation := []struct {
		input    string
		expected string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}
	for _, quot := range Quotation {
		output := QuotationMark(quot.input)

		if output != quot.expected {
			t.Errorf("\nGot: [%s] \n Expected: [%s]", output, quot.expected)
		}
	}
}

func TestPunctuationFunctionality(t *testing.T) {
	Punctuate := []struct {
		input    string
		expected string
	}{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
	}

	for _, sentence := range Punctuate {
		output := Punctuation(sentence.input)

		if output != sentence.expected {
			t.Errorf("\nGot: [%s] \nExpected: [%s] ", output, sentence.expected)
		}
	}
}

func TestVowelsFunctionality(t *testing.T) {
	vowels := []struct {
		input    string
		expected string
	}{
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
	}

	for _, vowel := range vowels {
		output := Vowels(vowel.input)

		if output != vowel.expected {
			t.Errorf("\nGot: [%s] \nExpected: [%s]", output, vowel.expected)
		}
	}
}

func TestProcessWordCommandsFunctionality(t *testing.T)  {

	commands := []struct {
		input string
		expexted string
	}{
		{"it (cap) was the best of times,",  "It was the best of times,"},
		{"it was the worst of times (up) ", "it was the worst of TIMES"},
		{"it was the age of wisdom, it was the age of foolishness (cap, 6) ", "it was the age of wisdom, It Was The Age Of Foolishness" },
		{"it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "it was the spring of hope, it was the winter of despair."},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
	}

	for _, line := range commands {
		output := ProcessWordCommands(strings.Fields(line.input))

		if output != line.expexted {
			t.Errorf("\nGot: [%s] \nExpected: [%s]", output, line.expexted)
		}
	}
	
}
